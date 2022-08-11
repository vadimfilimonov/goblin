package goblin

import (
	"reflect"
)

// IsBool checks if `arg` is classified as a bool.
func IsBool[T any](arg T) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.Bool
}

// IsBoolean checks if `arg` is classified as a bool.
func IsBoolean[T any](arg T) bool {
	return IsBool(arg)
}
