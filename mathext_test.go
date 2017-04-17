package mathext

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleHash() {
	fmt.Println(Hash("john"))
	fmt.Println(Hash("12345678"))
	fmt.Println(Hash("XXX_YYY_ZZZ"))
	fmt.Println(Hash("xxx_yyy_zzz"))
	// Output:
	// 6774539739450401392
	// -4898812128727250071
	// -8286756815414078424
	// -8259655320462518136
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
