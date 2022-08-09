package goblin

import (
	"reflect"
)

// IsString checks if 'arg' is classified as a `String`.
func IsString(arg interface{}) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.String
}
