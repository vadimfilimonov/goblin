package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(value int, index int, slice []int) bool
		want      []int
	}{
		{
			name:  "should return an empty slice",
			slice: []int{},
			predicate: func(value, index int, slice []int) bool {
				return true
			},
			want: []int{},
		},
		{
			name:  "should return an initial slice",
			slice: []int{1, 2, 3, 4},
			predicate: func(value, index int, slice []int) bool {
				return true
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name:  "should return first three numbers",
			slice: []int{1, 2, 3, 4, 5},
			predicate: func(value, index int, slice []int) bool {
				return value <= 3
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeWhile(tt.slice, tt.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}
