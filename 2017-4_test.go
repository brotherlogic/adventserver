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
		got := isValidPassword(c.in, false)
		if got != c.want {
			t.Errorf("Spec(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestPasphrase2(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, c := range cases {
		got := isValidPassword(c.in, true)
		if got != c.want {
			t.Errorf("Spec(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
