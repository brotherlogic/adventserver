package main

import "testing"

func TestDay1(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, c := range cases {
		got := ComputeFloor(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDayP2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, c := range cases {
		got := ComputeF1(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
