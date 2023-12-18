package math

import (
	"testing"
)

func TestPow(t *testing.T) {
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// zero
		{"0^0", args{0, 0}, 1},
		{"0^1", args{0, 1}, 0},
		{"0^2", args{0, 2}, 0},
		// one
		{"1^0", args{1, 0}, 1},
		{"1^1", args{1, 1}, 1},
		{"1^2", args{1, 2}, 1},
		// two
		{"2^0", args{2, 0}, 1},
		{"2^1", args{2, 1}, 2},
		{"2^2", args{2, 2}, 4},
		{"2^32", args{2, 32}, 4294967296},
		// large numbers
		{"55^0", args{55, 0}, 1},
		{"55^1", args{55, 1}, 55},
		{"55^2", args{55, 2}, 3025},
		{"55^3", args{55, 3}, 166375},
		// negative numbers
		{"-1^0", args{-1, 0}, 1},
		{"-1^1", args{-1, 1}, -1},
		{"-1^2", args{-1, 2}, 1},
		{"-1^3", args{-1, 3}, -1},
		// large negative numbers
		{"-55^0", args{-55, 0}, 1},
		{"-55^1", args{-55, 1}, -55},
		{"-55^2", args{-55, 2}, 3025},
		{"-55^3", args{-55, 3}, -166375},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
