package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropRight(t *testing.T) {
	assert.Equal(t, []int{1, 2}, DropRight([]int{1, 2, 3}, 1))
	assert.Equal(t, []int{1}, DropRight([]int{1, 2, 3}, 2))
	assert.Equal(t, []int{}, DropRight([]int{1, 2, 3}, 5))
	assert.Equal(t, []int{1, 2, 3}, DropRight([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{1, 2, 3}, DropRight([]int{1, 2, 3}, -1))
}
