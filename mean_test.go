package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	act, err := Mean([]float64{4, 2, 8, 6})
	assert.Equal(t, 5.0, act)
	assert.NoError(t, err)

	act, err = Mean([]float64{})
	assert.Equal(t, 0.0, act)
	assert.Error(t, err)
}
