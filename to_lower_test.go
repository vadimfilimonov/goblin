package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLower(t *testing.T) {
	assert.Equal(t, "--foo-bar--", ToLower("--Foo-Bar--"))
	assert.Equal(t, "foobar", ToLower("fooBar"))
	assert.Equal(t, "__foo_bar__", ToLower("__FOO_BAR__"))
}
