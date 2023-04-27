package goblin

// Tail gets all but the first element of slice.
func Tail[T any](slice []T) []T {
	if len(slice) <= 1 {
		emptySlice := make([]T, 0)
		return emptySlice
	}

	return slice[1:]
}
