package goblin

import "fmt"

// Last gets the last element of slice.
func Last[T any](collection []T) (T, error) {
	if len(collection) == 0 {
		var empty T
		return empty, fmt.Errorf("cannot get a last in an empty slice")
	}

	lastIndex := len(collection) - 1
	return collection[lastIndex], nil
}
