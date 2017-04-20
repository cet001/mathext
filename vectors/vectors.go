// Package vectors provides basic operations on vectors, as well as a sparse
// vector implementation.
package vectors

import (
	"github.com/cet001/mathext/ints"
	"math"
)

// Represents a sparse vector, where most of the elements are typically empty.
// For example, consider the following logical vector containing 10 elements:
//
//   v = [9 0 0 2 0 0 0 0 7 0]
//
// Only elements in ordinal positions 0, 3, and 8 contain non-zero values -- the
// remaining elements are "empty".  The following sparse vector is equivalent to
// the above vector:
//
//   sv := SparseVector{{0, 9}, {3, 2}, {8, 7}}
//
// Each element in the SparseVector is a Element object that specifies the
// element's value (Element.Value) and position (Element.Id) within the vector.
// Note that the SparseVector declaration above is a shorthand syntax; it can
// also be declared more formally like this:
//
//   sv := SparseVector{
//	   Element{Id: 0, Value: 9},
//	   Element{Id: 3, Value: 2},
//	   Element{Id: 8, Value: 7},
//   }
//
type SparseVector []Element

// An element within a SparseVector.  Id is the element's ordinal position within
// the vector, and Value is the element's value.
type Element struct {
	Id    int
	Value float64
}

// Calculates the dot product of two SparseVector objects.  Dot() assumes that
// v1 and v2 are in sorted order by Element.Id.
func Dot(v1, v2 SparseVector) float64 {
	var dotProduct float64 = 0.0
	lenV1, lenV2 := len(v1), len(v2)
	idx1, idx2 := 0, 0

	for {
		if idx1 == lenV1 || idx2 == lenV2 {
			break
		}

		e1, e2 := &v1[idx1], &v2[idx2]

		if e1.Id < e2.Id {
			idx1++
		} else if e2.Id < e1.Id {
			idx2++
		} else {
			dotProduct += (e1.Value * e2.Value)
			idx1++
			idx2++
		}
	}

	return dotProduct
}

// Calculates the Euclidean norm (a.k.a. L2-Norm) of the specified vector.
func Norm(vec SparseVector) float64 {
	sumOfSquares := 0.0
	for i := 0; i < len(vec); i++ {
		element := &vec[i]
		sumOfSquares += (element.Value * element.Value)
	}

	return math.Sqrt(sumOfSquares)
}

// Calculates a weighted mean for the specified values in []x and associated
// weights in []w.
//
// This function assumes that:
//    - x and w are are the same length
//    - all values in x and w are non-negative
func WeightedMean(x, w []float64) float64 {
	sumOfWeightedValues := 0.0
	sumOfWeights := 0.0
	for i, xVal := range x {
		sumOfWeightedValues += (xVal * w[i])
		sumOfWeights += w[i]
	}

	return sumOfWeightedValues / sumOfWeights
}

// Similar to the Unix 'uniq' command, this function removes all duplicates from
// a sorted array of int values.
func Uniq(sortedValues []int) []int {
	if sortedValues == nil {
		return []int{}
	}

	if len(sortedValues) <= 1 {
		return sortedValues
	}

	uniqueValues := make([]int, 0, len(sortedValues))
	uniqueValues = append(uniqueValues, sortedValues[0])

	for i := 1; i < len(sortedValues); i++ {
		if sortedValues[i] != sortedValues[i-1] {
			uniqueValues = append(uniqueValues, sortedValues[i])
		}
	}

	return uniqueValues
}

// Returns the intersection of 2 sorted sets.
//
// a and b are the sets to be intersected.
//
// target is an optional slice into which the intersecting elements from a and b
// are appended.  If target is nil, a new []int slice will be created and returned.
//
// WARNING: Unpredicatable results ensue if a or b contain duplicate elements or
// are not in ascending sorted order.
func Intersect(a, b, target []int) []int {
	lenA, lenB := len(a), len(b)

	var intersection []int
	if target == nil {
		intersection = make([]int, 0, ints.Min(lenA, lenB))
	} else {
		intersection = target[:0]
	}

	idx1, idx2 := 0, 0
	for {
		if idx1 == lenA || idx2 == lenB {
			break
		}

		aVal, bVal := a[idx1], b[idx2]

		if aVal < bVal {
			idx1++
		} else if bVal < aVal {
			idx2++
		} else {
			intersection = append(intersection, aVal)
			idx1++
			idx2++
		}
	}

	return intersection
}

// Returns the union of 2 sorted sets.
//
// a and b are the sets to be unioned.
//
// target is an optional slice into which the unique  elements from a and b are
// appended.  If this param is nil, a new []int slice will be created.
//
// WARNING: Unpredicatable results ensue if a or b contain duplicate elements or
// are not in ascending sorted order.

// Binary merge of sorted sets a and b.
// Unpredicatable results ensue if a or b contain duplicate elements or are not
// in ascending sorted order.
func Union(a, b, target []int) []int {
	lenA, lenB := len(a), len(b)

	var union []int
	if target == nil {
		union = make([]int, 0, ints.Max(lenA, lenB))
	} else {
		union = target[:0]
	}

	idx1, idx2 := 0, 0
	for {
		if idx1 == lenA {
			union = append(union, b[idx2:]...)
			break
		} else if idx2 == lenB {
			union = append(union, a[idx1:]...)
			break
		}

		aVal, bVal := a[idx1], b[idx2]

		if aVal < bVal {
			union = append(union, aVal)
			idx1++
		} else if bVal < aVal {
			union = append(union, bVal)
			idx2++
		} else {
			union = append(union, aVal)
			idx1++
			idx2++
		}
	}

	return union
}

// Sorts Elements by increasing Element.Id.
type ByElementId []Element

func (a ByElementId) Len() int           { return len(a) }
func (a ByElementId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByElementId) Less(i, j int) bool { return a[i].Id < a[j].Id }

// Sorts Elements by decreasing Element.Value.
type ByElementValueDesc []Element

func (a ByElementValueDesc) Len() int           { return len(a) }
func (a ByElementValueDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByElementValueDesc) Less(i, j int) bool { return a[i].Value > a[j].Value }
