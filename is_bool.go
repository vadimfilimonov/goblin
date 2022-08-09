package goblin

import (
	"reflect"
)

// IsBool checks if `arg` is classified as a bool.
func IsBool(arg interface{}) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.Bool
}

// IsBoolean checks if `arg` is classified as a bool.
func IsBoolean(arg interface{}) bool {
	return IsBool(arg)
}
