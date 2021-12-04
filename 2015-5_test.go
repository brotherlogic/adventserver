package main

import "testing"

func TestDay5P2a(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"xyxy", 2},
		{"aabcdefgaa", 2},
		{"aaa", 1},
		{"aaaa", 2},
	}

	for _, c := range cases {
		got := CountMaxNonOverlapping(c.in)
		if got != c.want {
			t.Errorf("NonOver(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay5P2b(t *testing.T) {

	cases := []struct {
		in   string
		want bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
	}

	for _, c := range cases {
		got := RepeatWithMiddle(c.in)
		if got != c.want {
			t.Errorf("NonOver(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestDay5P2(t *testing.T) {

	cases := []struct {
		in   string
		want bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"xyxaaabgaa", true},
	}

	for _, c := range cases {
		got := IsNiceAlso(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestDay5P1(t *testing.T) {

	cases := []struct {
		in   string
		want bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, c := range cases {
		got := IsNice(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
