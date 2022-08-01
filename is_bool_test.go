package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBool(t *testing.T) {
	assert.Equal(t, true, IsBool(true))
	assert.Equal(t, true, IsBool(false))

	assert.Equal(t, false, IsBool("abc"))
	assert.Equal(t, false, IsBool([]int{1, 2, 3}))
	assert.Equal(t, false, IsBool([]string{"1", "2", "3"}))
	assert.Equal(t, false, IsBool(map[string]bool{"foo": true}))
	assert.Equal(t, false, IsBool(42))
}
