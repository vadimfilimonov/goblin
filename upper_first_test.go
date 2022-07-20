package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperFirst(t *testing.T) {
	assert.Equal(t, UpperFirst("fred"), "Fred")
	assert.Equal(t, UpperFirst("FRED"), "FRED")
	assert.Equal(t, UpperFirst(""), "")
}
