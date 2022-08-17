package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "a~b~c", Join([]string{"a", "b", "c"}, "~"))
	assert.Equal(t, "", Join([]string{}, ","))
}
