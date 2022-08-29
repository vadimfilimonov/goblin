package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	assert.Equal(t, []int{1}, Take([]int{1, 2, 3}, 1))
	assert.Equal(t, []int{1, 2}, Take([]int{1, 2, 3}, 2))
	assert.Equal(t, []int{1, 2, 3}, Take([]int{1, 2, 3}, 5))
	assert.Equal(t, []int{}, Take([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{}, Take([]int{1, 2, 3}, -1))
}
