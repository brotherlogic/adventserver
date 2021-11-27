package main

import "testing"

func Test2016Day1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
		{"R2, R2, R2, R2", 0},
		{"R2, R2, R2, R2, R2, R2, R2, R2, R2, R2, R2, R2", 0},
		{"L2, L2, L2, L2, L2, L2, L2, L2, L2, L2, L2, L2", 0},
	}

	for _, c := range cases {
		got := runCoords(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
