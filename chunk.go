package goblin

// Chunk creates a slice of elements split into groups the length of `size`. If `slice` can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 || len(slice) == 0 {
		return [][]T{}
	}

	chunkedSlice := make([][]T, 0, size)

	for i := 0; i < len(slice); i += size {
		chunkEndIndex := i + size
		if chunkEndIndex > len(slice) {
			chunkEndIndex = len(slice)
		}

		chunk := make([]T, 0, size)
		for j := i; j < chunkEndIndex; j++ {
			chunk = append(chunk, slice[j])
		}
		chunkedSlice = append(chunkedSlice, chunk)
	}

	return chunkedSlice
}
