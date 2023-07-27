package goblin

import (
	"fmt"
	"math"
	"reflect"
)

// Compact creates a slice with all falsey values removed. The values `false, 0, ""`, and `NaN` are falsey.
func Compact[T comparable](slice []T) ([]T, error) {
	compactedSlice := make([]T, 0)

	for i := 0; i < len(slice); i++ {
		item := slice[i]

		switch any(item).(type) {
		case string:
			{
				isFalsey := reflect.ValueOf(item).String() == ""
				if !isFalsey {
					compactedSlice = append(compactedSlice, item)
				}
			}
		case int, int8, int16, int32, int64:
			{
				isFalsey := reflect.ValueOf(item).Int() == 0
				if !isFalsey {
					compactedSlice = append(compactedSlice, item)
				}
			}
		case uint, uint8, uint16, uint32, uint64:
			{
				isFalsey := reflect.ValueOf(item).Uint() == 0
				if !isFalsey {
					compactedSlice = append(compactedSlice, item)
				}
			}
		case float32, float64:
			{
				number := reflect.ValueOf(item).Float()
				isFalsey := number == 0 || math.IsNaN(number)
				if !isFalsey {
					compactedSlice = append(compactedSlice, item)
				}
			}
		case bool:
			{
				isFalsey := !reflect.ValueOf(item).Bool()
				if !isFalsey {
					compactedSlice = append(compactedSlice, item)
				}
			}
		default:
			{
				return nil, fmt.Errorf("unsupported type of slice's elements")
			}
		}
	}

	return compactedSlice, nil
}
