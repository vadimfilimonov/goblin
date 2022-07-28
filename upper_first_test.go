package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperFirst(t *testing.T) {
	assert.Equal(t, "Fred", UpperFirst("fred"))
	assert.Equal(t, "FRED", UpperFirst("FRED"))
	assert.Equal(t, "", UpperFirst(""))
}
