package goblin

// Gets all but the last element of `slice`.
func Initial[T any](slice []T) []T {
	return DropRight(slice, 1)
}
