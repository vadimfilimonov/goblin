package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	act, err := Max([]float64{4, 2, 8, 6})
	assert.Equal(t, act, float64(8))
	assert.NoError(t, err)

	act, err = Max([]float64{})
	assert.Equal(t, act, float64(0))
	assert.Error(t, err)
}
