package mathext

import (
	"math"
)

// Convenience coefficient for converting degrees to radians.
const Deg2rad float64 = math.Pi / 180

// Hashes a string into an int.
func Hash(s string) int {
	// NOTE: This implementation based on Java's String.hashCode()

	h := 1125899906842597 // prime
	len := len(s)

	for i := 0; i < len; i++ {
		h = 31*h + int(s[i])
	}

	return h
}
