package goblin

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 1.5, Divide(6, 4))
	assert.Equal(t, true, math.IsNaN(Divide(0, 0)))
	assert.Equal(t, math.Inf(-1), Divide(-42, 0))
}
