package goblin

// Take creates a slice of `slice` with `n` elements taken from the beginning.
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}

	if n > len(slice) {
		return slice
	}

	return slice[:n]
}
