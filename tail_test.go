package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTail(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Tail([]int{1, 2, 3}))
	assert.Equal(t, []string{}, Tail([]string{}))
}
