package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubFalse(t *testing.T) {
	act := StubFalse()
	assert.Equal(t, act, false)
}
