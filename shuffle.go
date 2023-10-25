package goblin

import (
	"crypto/rand"
	"math/big"
)

// Shuffle creates a slice of shuffled values.
func Shuffle[T any](slice []T) ([]T, error) {
	shuffledSlice := make([]T, len(slice))
	copy(shuffledSlice, slice)

	for index := len(shuffledSlice) - 1; index > 0; index-- {
		randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(index+1)))
		if err != nil {
			return shuffledSlice, err
		}
		randomIndex := randomNumber.Int64()
		shuffledSlice[index], shuffledSlice[randomIndex] = shuffledSlice[randomIndex], shuffledSlice[index]
	}

	return shuffledSlice, nil
}
