package main

import "testing"

func TestAdd(t *testing.T) {
		tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 3, want: 5},
	}
	for i, tt := range tests {
		if got := add(tt.a, tt.b); got != tt.want {
			t.Errorf("#%d: Add(%v, %v) = %v, want %v", i, tt.a, tt.b, got, tt.want)
		}
	}
}
