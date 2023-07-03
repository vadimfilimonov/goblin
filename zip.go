package goblin

// Zip creates a slice of grouped elements, the first of which contains the first elements of the given arrays, the second of which contains the second elements of the given arrays, and so on.
func Zip[T any](slice [][]T) [][]T {
	if len(slice) == 0 {
		return [][]T{}
	}

	zippedSliceLength := len(slice[0])
	zippedSlice := make([][]T, zippedSliceLength)

	for i := 0; i < zippedSliceLength; i += 1 {
		zippedSlice[i] = make([]T, len(slice))

		for j := 0; j < len(slice); j += 1 {
			zippedSlice[i][j] = slice[j][i]
		}
	}

	return zippedSlice
}
