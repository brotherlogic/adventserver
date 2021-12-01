package main

import "testing"

func Test2017Day3P1(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, c := range cases {
		got := computeSpiral(c.in)
		if int(got) != c.want {
			t.Errorf("Spec(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Test2017Day3P2(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{120, 122},
	}

	for _, c := range cases {
		got := buildSpiral(c.in)
		if int(got) != c.want {
			t.Errorf("Spec(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
