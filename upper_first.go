package goblin

import (
	"strings"
)

// Converts the first character of `string` to upper case.
func UpperFirst(text string) string {
	return strings.Title(text)
}
