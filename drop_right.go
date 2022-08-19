package goblin

// DropRight creates a slice of `array` with `n` elements dropped from the end.
func DropRight[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}

	if n > len(slice) {
		return []T{}
	}

	end := len(slice) - n

	return slice[:end]
}
