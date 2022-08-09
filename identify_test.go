package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	assert.Equal(t, true, Identity(true))
	assert.Equal(t, "abc", Identity("abc"))
	assert.Equal(t, 42, Identity(42))
	assert.Equal(t, []int{1, 2, 3}, Identity([]int{1, 2, 3}))
	assert.Equal(t, map[string]bool{"foo": true}, Identity(map[string]bool{"foo": true}))
}
