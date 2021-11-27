package main

import "testing"

func Test2017Day1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, c := range cases {
		got := computeDigs(c.in)
		if int(got) != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
