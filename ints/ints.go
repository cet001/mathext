package ints

// Returns the lesser of a and b.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns the greater of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
