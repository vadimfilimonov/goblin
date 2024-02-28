package goblin

import (
	"reflect"
	"testing"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "success",
			slice: []int{2, 1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "empty",
			slice: []int{2, 1, 2},
			want:  []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uniq(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
