package goblin

import (
	"math/rand"
	"time"
)

// Shuffle creates a slice of shuffled values.
func Shuffle[T any](slice []T) []T {
	shuffledSlice := make([]T, len(slice))
	copy(shuffledSlice, slice)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffledSlice), func(i, j int) {
		shuffledSlice[i], shuffledSlice[j] = shuffledSlice[j], shuffledSlice[i]
	})

	return shuffledSlice
}
