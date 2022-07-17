package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(6, 4)
	assert.Equal(t, sum, 10)
}
