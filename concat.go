package goblin

func Concat[T any](slice []T, values ...[]T) []T {
	concatedSlice := make([]T, len(slice))
	copy(concatedSlice, slice)

	if len(values) == 0 {
		return concatedSlice
	}

	for _, value := range values {
		concatedSlice = append(concatedSlice, value...)
	}

	return concatedSlice
}
