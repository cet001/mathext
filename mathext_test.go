package mathext

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleHash() {
	fmt.Println(Hash("john"))
	fmt.Println(Hash("John"))
	fmt.Println(Hash("12345678"))
	// Output:
	// 6774539739450401392
	// 6774539739449448080
	// -4898812128727250071
}

// This is just a sanity-check.
func TestHash(t *testing.T) {
	values := []string{
		"", "a", "b", "c", "A", "B", "C", "cat", "CAT",
		"aaaaaaaaaaaaaaaa", "???????????????????????",
		"1", " 1", "  1",
	}

	uniqueHashValues := map[int]bool{}
	for _, value := range values {
		uniqueHashValues[Hash(value)] = true
	}

	assert.Equal(t, len(values), len(uniqueHashValues))
}

func BenchmarkHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash("abcdefgABCDEFG012345") // 20 character rstring
	}
}
