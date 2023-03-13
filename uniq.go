package goblin

// Uniq creates a duplicate-free version of an array. The order of result values is determined by the order they occur in the array.
func Uniq[T comparable](slice []T) []T {
	uniqSlice := make([]T, 0)

	for i := 0; i < len(slice); i++ {
		currentValue := slice[i]

		if !Includes(uniqSlice, currentValue, 0) {
			uniqSlice = append(uniqSlice, currentValue)
		}
	}

	return uniqSlice
}
