package goblin

import (
	"reflect"
)

// IsString checks if 'arg' is classified as a `String`.
func IsString[T any](arg T) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.String
}
