package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGt(t *testing.T) {
	assert.Equal(t, true, Gt(3, 1))
	assert.Equal(t, false, Gt(3, 3))
	assert.Equal(t, false, Gt(1, 3))
}
