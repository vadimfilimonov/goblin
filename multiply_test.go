package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, 24.0, Multiply(6, 4))
	assert.Equal(t, 0.0, Multiply(42, 0))
}
