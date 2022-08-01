package goblin

import (
	"reflect"
)

// Checks if `arg` is classified as a bool.
func IsBool(arg interface{}) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.Bool
}

// Checks if `arg` is classified as a bool.
func IsBoolean(arg interface{}) bool {
	return IsBool(arg)
}
