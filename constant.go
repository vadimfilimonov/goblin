package goblin

// Constant creates a function that returns `value`.
func Constant[T any](arg T) func() T {
	return func() T {
		return arg
	}
}
