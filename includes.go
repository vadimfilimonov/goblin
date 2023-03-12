package goblin

import (
	"math"
)

// Includes checks if `value` is in `collection`.
func Includes[T comparable](collection []T, value T, fromIndex int) bool {
	if len(collection) == 0 {
		return false
	}

	if fromIndex < 0 {
		fromIndex = len(collection) - int(math.Abs(float64(fromIndex)))
	}

	if fromIndex < 0 {
		fromIndex = 0
	}

	for i := fromIndex; i < len(collection); i++ {
		currentValue := collection[i]

		if currentValue == value {
			return true
		}
	}

	return false
}
