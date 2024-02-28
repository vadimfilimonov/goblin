package goblin

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		augend float64
		addend float64
		want   float64
	}{
		{
			name:   "Integers",
			augend: 6,
			addend: 4,
			want:   10.0,
		},
		{
			name:   "Fractional numbers",
			augend: 6.5,
			addend: 4.5,
			want:   11.0,
		},
		{
			name:   "Fractional and integer numbers",
			augend: 6.5,
			addend: 4,
			want:   10.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.augend, tt.addend); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
