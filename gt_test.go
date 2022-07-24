package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGt(t *testing.T) {
	assert.Equal(t, Gt(float64(3), float64(1)), true)
	assert.Equal(t, Gt(float64(3), float64(3)), false)
	assert.Equal(t, Gt(float64(1), float64(3)), false)
}
