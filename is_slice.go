package goblin

import (
	"reflect"
)

// IsSlice checks if 'arg' is classified as a 'Slice'
func IsSlice[T any](arg T) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.Slice
}
