package goblin

// Drop creates a slice of `slice` with `n` elements dropped from the beginning.
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}

	if n > len(slice) {
		return []T{}
	}

	return slice[n:]
}
