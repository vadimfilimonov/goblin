package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	assert.Equal(t, []int{}, Uniq([]int{}))
	assert.Equal(t, []int{2, 1}, Uniq([]int{2, 1, 2}))
}
