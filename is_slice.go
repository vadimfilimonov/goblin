package goblin

import (
	"reflect"
)

// IsSlice checks if 'arg' is classified as a 'Slice'
func IsSlice(arg interface{}) bool {
	val := reflect.ValueOf(arg)

	return val.Kind() == reflect.Slice
}
