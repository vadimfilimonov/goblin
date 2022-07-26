package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLt(t *testing.T) {
	assert.Equal(t, Lt(float64(1), float64(3)), true)
	assert.Equal(t, Lt(float64(3), float64(3)), false)
	assert.Equal(t, Lt(float64(3), float64(1)), false)
}
