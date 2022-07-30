package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSlice(t *testing.T) {
	assert.Equal(t, true, IsSlice([]int{1, 2, 3}))
	assert.Equal(t, true, IsSlice([]string{"1", "2", "3"}))

	assert.Equal(t, false, IsSlice(map[string]bool{"foo": true}))
	assert.Equal(t, false, IsSlice("abc"))
	assert.Equal(t, false, IsSlice(42))
}
