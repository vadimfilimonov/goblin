package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvert(t *testing.T) {
	actual := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := map[int]string{1: "a", 2: "b", 3: "c"}

	assert.Equal(t, expected, Invert(actual))
}
