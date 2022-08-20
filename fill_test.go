package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	assert.Equal(t, []int{7, 7, 7}, Fill([]int{1, 2, 3}, 7, 0, 3))
	assert.Equal(t, []string{"a", "*", "*", "d"}, Fill([]string{"a", "b", "c", "d"}, "*", 1, 3))

	slice := []bool{true, true}
	assert.Equal(t, []bool{false, false}, Fill(slice, false, 0, 2))
	assert.Equal(t, []bool{true, true}, slice)
}
