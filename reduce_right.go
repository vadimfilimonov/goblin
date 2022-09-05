package goblin

// ReduceRight is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight[T any, K any](slice []T, callback func(accumulator K, value T, index int, slice []T) K, initAccumulator K) K {
	accumulator := initAccumulator

	for index := len(slice) - 1; index >= 0; index -= 1 {
		value := slice[index]
		accumulator = callback(accumulator, value, index, slice)
	}

	return accumulator
}
