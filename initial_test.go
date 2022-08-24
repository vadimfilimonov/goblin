package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitial(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Initial([]int{1, 2, 3}))
}
