package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	arg := map[string]bool{"foo": true, "bar": false}
	myFunc := Constant(arg)

	assert.Equal(t, arg, myFunc())
	assert.Equal(t, "Run", Constant("Run")())
}
