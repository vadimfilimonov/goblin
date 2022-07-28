package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 10.0, Add(6, 4))
	assert.Equal(t, 11.0, Add(6.5, 4.5))
	assert.Equal(t, 10.5, Add(6.5, 4))
}
