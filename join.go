package goblin

import (
	"strings"
)

// Join converts all elements in slice into a string separated by separator.
func Join(slice []string, separator string) string {
	// TODO: support numbers in slice
	return strings.Join(slice, separator)
}
