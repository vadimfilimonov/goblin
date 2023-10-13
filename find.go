package goblin

// Find iterates over elements of slice, returning the first element predicate returns truthy for. The predicate is invoked with three arguments: (value, index, collection).
func Find[T any](slice []T, callback func(item T, index int, slice []T) bool) (findedValue T, ok bool) {
	for index, value := range slice {
		if callback(value, index, slice) {
			return value, true
		}
	}

	var emptyValue T
	return emptyValue, false
}
