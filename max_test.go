package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	act, err := Max([]float64{4, 2, 8, 6})
	assert.Equal(t, 8.0, act)
	assert.NoError(t, err)

	act, err = Max([]float64{})
	assert.Equal(t, 0.0, act)
	assert.Error(t, err)
}
