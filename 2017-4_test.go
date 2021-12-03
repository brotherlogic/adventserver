package main

import "testing"

func TestPasphrase(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, c := range cases {
		got := isValidPassword(c.in)
		if got != c.want {
			t.Errorf("Spec(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
