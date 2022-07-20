package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubString(t *testing.T) {
	act := StubString()
	assert.Equal(t, act, "")
}
