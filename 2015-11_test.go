package main

import "testing"

func TestIncrementString(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"xx", "xy"},
		{"xy", "xz"},
		{"xz", "ya"},
		{"ya", "yb"},
		{"yb", "yc"},
	}

	for _, c := range cases {
		out := incrementString(c.in)
		if out != c.out {
			t.Fatalf("Bad trans %v => %v (should have been %v)", c.in, out, c.out)
		}
	}
}

func TestBadPass(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdefgh", false},
		{"abcdffaa", true},
		{"ghijklmn", false},
		{"ghjaabcc", true},
	}

	for _, c := range cases {
		out := isValidSantaPassword(c.in)
		if out != c.out {
			t.Fatalf("Bad trans %v => %v (should have been %v)", c.in, out, c.out)
		}
	}
}

func TestFindNext(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, c := range cases {
		out := findNextPassword(c.in)
		if out != c.out {
			t.Fatalf("Bad trans %v => %v (should have been %v)", c.in, out, c.out)
		}
	}
}
