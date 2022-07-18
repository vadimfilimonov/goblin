package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubTrue(t *testing.T) {
	act := StubTrue()
	assert.Equal(t, act, true)
}
