package goblin

// TakeWhile creates a slice of `slice` with elements taken from the beginning. Elements are taken until `predicate` returns falsey. The predicate is invoked with three arguments: _(value, index, array)_.
func TakeWhile[T any](slice []T, predicate func(value T, index int, slice []T) bool) []T {
	if len(slice) == 0 || slice == nil {
		return slice
	}

	result := make([]T, 0)

	for i, v := range slice {
		if !predicate(v, i, slice) {
			return result
		}

		result = append(result, v)
	}

	return result
}
