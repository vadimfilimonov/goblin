package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(float64(6), float64(4))
	assert.Equal(t, sum, float64(10))
}
