package goblin

// TakeRight creates a slice of `slice` with `n` elements taken from the end.
func TakeRight[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}

	if n > len(slice) {
		return slice
	}

	index := len(slice) - n

	return slice[index:]
}
