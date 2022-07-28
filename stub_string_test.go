package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubString(t *testing.T) {
	assert.Equal(t, "", StubString())
}
