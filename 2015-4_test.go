package main

import "testing"

func TestDay4P1(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, c := range cases {
		got := solveHash(c.in, 5)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
