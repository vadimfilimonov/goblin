package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubTrue(t *testing.T) {
	assert.Equal(t, true, StubTrue())
}
