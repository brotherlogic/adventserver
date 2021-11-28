package main

import "testing"

func TestDay2P2(t *testing.T) {
	cases := []struct {
		in   string
		want int32
	}{
		{"2x3x4", int32(34)},
		{"1x1x10", int32(14)},
	}

	for _, c := range cases {
		got := computeAmountOfRibbon(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay2(t *testing.T) {

	cases := []struct {
		in   string
		want int32
	}{
		{"2x3x4", int32(58)},
		{"1x1x10", int32(43)},
	}

	for _, c := range cases {
		got := computeAmountOfPaper(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
