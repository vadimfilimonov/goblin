package goblin

// Fills elements of `slice` with `value` from `start` up to, but not including, `end`.
func Fill[T any](slice []T, value T, start int, end int) []T {
	clonedSlice := []T{}
	clonedSlice = append(clonedSlice, slice...)

	for index := start; index < end; index += 1 {
		clonedSlice[index] = value
	}

	return clonedSlice
}
