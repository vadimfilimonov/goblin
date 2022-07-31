package goblin

import (
	"reflect"
)

// Checks if 'arg' is classified as a `String`.
func IsString(arg interface{}) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.String
}
