package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	difference := Subtract(6, 4)
	assert.Equal(t, difference, 2)

	difference = Subtract(4, 6)
	assert.Equal(t, difference, -2)
}
