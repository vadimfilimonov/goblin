package goblin

// ForEachRight method is like ForEach except that it iterates over elements of collection from right to left.
func ForEachRight[T any](collection []T, iteratee func(value T)) {
	for i := len(collection) - 1; i >= 0; i-- {
		iteratee(collection[i])
	}
}
