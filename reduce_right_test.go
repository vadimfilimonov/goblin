package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceRight(t *testing.T) {
	input := [][]int{{0, 1}, {2, 3}, {4, 5}}
	actual := ReduceRight(input, func(accumulator []int, value []int, index int, slice [][]int) []int {
		return append(accumulator, value...)
	}, []int{})
	expected := []int{4, 5, 2, 3, 0, 1}

	assert.Equal(t, expected, actual)
}
