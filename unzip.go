package goblin

// Unzip is like zip except that it accepts an array of grouped elements and creates an array regrouping the elements to their pre-zip configuration.
func Unzip[T any](slice [][]T) [][]T {
	return Zip(slice...)
}
