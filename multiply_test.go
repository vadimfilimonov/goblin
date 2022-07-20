package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, Multiply(float64(6), float64(4)), float64(24))
	assert.Equal(t, Multiply(float64(42), float64(0)), float64(0))
}
