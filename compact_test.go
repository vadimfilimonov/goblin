package goblin

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompact(t *testing.T) {
	actual, err := Compact([]string{"a", "", "b"})
	assert.Equal(t, []string{"a", "b"}, actual)
	assert.NoError(t, err)

	actual2, err := Compact([]int64{-1, 0, 1, 2})
	assert.Equal(t, []int64{-1, 1, 2}, actual2)
	assert.NoError(t, err)

	actual3, err := Compact([]float64{-1, 0, 1, math.NaN(), 2})
	assert.Equal(t, []float64{-1, 1, 2}, actual3)
	assert.NoError(t, err)

	actual4, err := Compact([]bool{true, false, true})
	assert.Equal(t, []bool{true, true}, actual4)
	assert.NoError(t, err)
}
