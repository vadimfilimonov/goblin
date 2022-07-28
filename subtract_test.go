package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	assert.Equal(t, 2.0, Subtract(6, 4))
	assert.Equal(t, -2.0, Subtract(4, 6))
}
