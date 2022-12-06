package main

import "testing"

func Test2022_6_1_Main(t *testing.T) {
	tests := []struct {
		in   string
		char int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range tests {
		first := findMarker(test.in)
		if first != test.char {
			t.Errorf("Bad marker %v: %v (%v)", test.in, first, test.char)
		}
	}
}
