package goblin

// Intersection an slice of unique values that are included in all given slices. The order and references of result values are determined by the first slice.
func Intersection[T comparable](slices ...[]T) []T {
	firstSlice := Uniq(slices[0])
	restSlices := slices[1:]

	if len(restSlices) == 0 {
		return firstSlice
	}

	result := make([]T, 0)

	// TODO: Refactor to map
	for _, item := range firstSlice {
		shouldAppendItemToResult := true

		for _, slice := range restSlices {
			if !Includes(slice, item, 0) {
				shouldAppendItemToResult = false
				break
			}
		}

		if shouldAppendItemToResult {
			result = append(result, item)
		}
	}

	return result
}
