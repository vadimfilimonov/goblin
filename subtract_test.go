package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	difference := Subtract(float64(6), float64(4))
	assert.Equal(t, difference, float64(2))

	difference = Subtract(float64(4), float64(6))
	assert.Equal(t, difference, float64(-2))
}
