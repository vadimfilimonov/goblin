package goblin

// Invert creates an map composed of the inverted keys and values of map.
// If map contains duplicate values, values overwrite property assignments in random order.
func Invert[K comparable, V comparable](input map[K]V) map[V]K {
	invertedMap := map[V]K{}

	for key, value := range input {
		invertedMap[value] = key
	}

	return invertedMap
}
