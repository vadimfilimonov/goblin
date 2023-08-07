package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	assert.Equal(t, []int{42}, Intersection([]int{42}))
	assert.Equal(t, []int{2}, Intersection([]int{2, 1}, []int{2, 3}))
	assert.Equal(t, []int{2}, Intersection([]int{2, 1, 2}, []int{2, 3}))
	assert.Equal(t, []int{3}, Intersection([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5}))
	assert.Equal(t, []int{}, Intersection([]int{1, 2, 3}, []int{}))
}
