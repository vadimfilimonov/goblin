package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeRight(t *testing.T) {
	assert.Equal(t, []int{3}, TakeRight([]int{1, 2, 3}, 1))
	assert.Equal(t, []int{2, 3}, TakeRight([]int{1, 2, 3}, 2))
	assert.Equal(t, []int{1, 2, 3}, TakeRight([]int{1, 2, 3}, 5))
	assert.Equal(t, []int{}, TakeRight([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{}, TakeRight([]int{1, 2, 3}, -1))
}
