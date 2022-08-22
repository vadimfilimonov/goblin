package goblin

// Map creates an slice of values by running each element in `slice` thru `iteratee`. The iteratee is invoked with three arguments: (value, index, slice).
func Map[T any, K any](slice []T, callback func(item T, index int, slice []T) K) []K {
	result := make([]K, len(slice))

	for index, value := range slice {
		result[index] = callback(value, index, slice)
	}

	return result
}
