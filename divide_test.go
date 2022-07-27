package goblin

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, float64(1.5), Divide(float64(6), float64(4)))
	assert.Equal(t, true, math.IsNaN(Divide(float64(0), float64(0))))
	assert.Equal(t, math.Inf(-1), Divide(float64(-42), float64(0)))
}
