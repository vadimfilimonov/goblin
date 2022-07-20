package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, Multiply(6, 4), 24)
	assert.Equal(t, Multiply(42, 0), 0)
}
