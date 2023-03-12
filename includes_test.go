package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludes(t *testing.T) {
	assert.Equal(t, true, Includes([]int{1, 2, 3}, 1, 0))
	assert.Equal(t, false, Includes([]int{1, 2, 3}, 1, 2))
	assert.Equal(t, false, Includes([]int{1, 2, 3, 4, 5}, 3, -2))
	assert.Equal(t, true, Includes([]int{1, 2, 3, 4, 5}, 3, -3))
	assert.Equal(t, true, Includes([]int{1, 2, 3, 4, 5}, 3, -10))
	assert.Equal(t, false, Includes([]int{}, 1, 0))
	assert.Equal(t, false, Includes([]int{}, 1, -2))
}
