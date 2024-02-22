package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat_Success(t *testing.T) {
	baseSlice := []int{1, 2, 3}
	secondSlice := []int{4, 5, 6}
	thirdSlice := []int{7, 8, 9}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := Concat(baseSlice, secondSlice, thirdSlice)

	assert.Equal(t, expected, actual)
}

func TestConcat_Empty(t *testing.T) {
	slice := []int{1, 2, 3}
	expected := []int{1, 2, 3}

	actual := Concat(slice)
	slice[0] = 42

	assert.Equal(t, expected, actual)
}
