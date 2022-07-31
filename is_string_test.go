package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsString(t *testing.T) {
	assert.Equal(t, true, IsString("abc"))
	assert.Equal(t, true, IsString(""))

	assert.Equal(t, false, IsString([]int{1, 2, 3}))
	assert.Equal(t, false, IsString([]string{"1", "2", "3"}))
	assert.Equal(t, false, IsString(map[string]bool{"foo": true}))
	assert.Equal(t, false, IsString(42))
	assert.Equal(t, false, IsString(true))
}
