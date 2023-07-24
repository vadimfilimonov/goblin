package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	input := []int{1, 2, 3, 4}
	actual := Shuffle(input)

	assert.ElementsMatch(t, actual, input)
}
