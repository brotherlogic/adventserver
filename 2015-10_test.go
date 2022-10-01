package main

import "testing"

func TestDay201510P1(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, c := range cases {
		out := lookAndSay(c.in)
		if out != c.out {
			t.Fatalf("Bad trans %v => %v (should have been %v)", c.in, out, c.out)
		}
	}
}
