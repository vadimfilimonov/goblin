package goblin

// Reduces `collection` to a value which is the accumulated result of running each element in `collection` thru `iteratee`, where each successive invocation is supplied the return value of the previous.
// The iteratee is invoked with four arguments: (accumulator, value, index, collection).
func Reduce[T any, K any](slice []T, callback func(accumulator K, value T, index int, slice []T) K, initAccumulator K) K {
	accumulator := initAccumulator

	for index, value := range slice {
		accumulator = callback(accumulator, value, index, slice)
	}

	return accumulator
}
