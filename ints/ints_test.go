package mathext

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1, Min(2, 1))

	assert.Equal(t, -2, Min(-1, -2))
	assert.Equal(t, -2, Min(-2, -1))

	assert.Equal(t, -1, Min(-1, 2))
	assert.Equal(t, -1, Min(2, -1))

	assert.Equal(t, 3, Min(3, 3))
	assert.Equal(t, -3, Min(-3, -3))

	assert.Equal(t, 0, Min(99, 0))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 2, Max(2, 1))

	assert.Equal(t, -1, Max(-1, -2))
	assert.Equal(t, -1, Max(-2, -1))

	assert.Equal(t, 2, Max(-1, 2))
	assert.Equal(t, 2, Max(2, -1))

	assert.Equal(t, 3, Max(3, 3))
	assert.Equal(t, -3, Max(-3, -3))

	assert.Equal(t, 99, Max(99, 0))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, 1, Abs(-1))
	assert.Equal(t, 0, Abs(0))
}
