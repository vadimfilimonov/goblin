package goblin

import "fmt"

// Head gets the first element of slice.
func Head[T any](collection []T) (T, error) {
	if len(collection) == 0 {
		var empty T
		return empty, fmt.Errorf("cannot get a head in an empty slice")
	}

	return collection[0], nil
}
