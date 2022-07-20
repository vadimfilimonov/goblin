package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	act, err := Min([]float64{4, 2, 8, 6})
	assert.Equal(t, act, float64(2))
	assert.NoError(t, err)

	act, err = Min([]float64{})
	assert.Equal(t, act, float64(0))
	assert.Error(t, err)
}
