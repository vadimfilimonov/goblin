package goblin

// Filter iterates over elements of `slice`, returning an slice of all elements `predicate` returns truthy for. The predicate is invoked with three arguments: (value, index, slice).
func Filter[T any](slice []T, callback func(item T, index int, slice []T) bool) []T {
	result := make([]T, 0)

	for index, value := range slice {
		if callback(value, index, slice) {
			result = append(result, value)
		}
	}

	return result
}
