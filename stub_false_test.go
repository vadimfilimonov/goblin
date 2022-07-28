package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubFalse(t *testing.T) {
	assert.Equal(t, false, StubFalse())
}
