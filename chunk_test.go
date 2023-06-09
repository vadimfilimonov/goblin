package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	assert.Equal(t, [][]string{}, Chunk([]string{"a", "b", "c", "d"}, 0))
	assert.Equal(t, [][]string{}, Chunk([]string{"a", "b", "c", "d"}, -1))
	assert.Equal(t, [][]string{}, Chunk([]string{}, 5))

	assert.Equal(t, [][]string{{"a", "b"}, {"c", "d"}}, Chunk([]string{"a", "b", "c", "d"}, 2))
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d"}}, Chunk([]string{"a", "b", "c", "d"}, 3))
}
