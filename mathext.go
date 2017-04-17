package mathext

// Hashes a string into an int.  This operation is useful if you want to convert
// a weighted string vector into a SparseVector. I.e. each string s[i] in a
// weighted vector V=[s0w0, s1w1, ...] can be converted into an int.
func Hash(s string) int {
	// NOTE: this implementation based on Java's String.hashCode()

	h := 1125899906842597 // prime
	len := len(s)

	for i := 0; i < len; i++ {
		h = 31*h + int(s[i])
	}

	return h
}
