package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEachRight(t *testing.T) {
	input := []int64{1, 2, 3}

	actual := make([]int64, 0, len(input))
	ForEachRight(input, func(value int64) {
		actual = append(actual, value*2)
	})

	assert.Equal(t, []int64{6, 4, 2}, actual)
}
