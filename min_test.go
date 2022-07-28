package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	act, err := Min([]float64{4, 2, 8, 6})
	assert.Equal(t, 2.0, act)
	assert.NoError(t, err)

	act, err = Min([]float64{})
	assert.Equal(t, 0.0, act)
	assert.Error(t, err)
}
