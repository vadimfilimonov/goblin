package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	actual1, err1 := Head([]int{1, 2, 3})
	assert.Equal(t, 1, actual1)
	assert.NoError(t, err1)

	actual2, err2 := Head([]string{})
	assert.Equal(t, "", actual2)
	assert.Error(t, err2)
}
