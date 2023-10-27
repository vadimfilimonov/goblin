package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	actual1, err1 := Last([]int{1, 2, 3})
	assert.Equal(t, 3, actual1)
	assert.NoError(t, err1)

	actual2, err2 := Last([]string{})
	assert.Equal(t, "", actual2)
	assert.Error(t, err2)
}
