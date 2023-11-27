package goblin

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(value T)) {
	for i := 0; i < len(collection); i++ {
		iteratee(collection[i])
	}
}
