package goblin

// Identity returns the first argument it receives.
func Identity[T any](arg T) T {
	return arg
}
