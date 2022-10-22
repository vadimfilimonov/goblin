package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUpper(t *testing.T) {
	assert.Equal(t, "--FOO-BAR--", ToUpper("--foo-bar--"))
	assert.Equal(t, "FOOBAR", ToUpper("fooBar"))
	assert.Equal(t, "__FOO_BAR__", ToUpper("__foo_bar__"))
}
