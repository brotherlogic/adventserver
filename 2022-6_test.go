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
		first := findMarker(test.in, 4)
		if first != test.char {
			t.Errorf("Bad marker %v: %v (%v)", test.in, first, test.char)
		}
	}
}

func Test2022_6_2_Main(t *testing.T) {
	tests := []struct {
		in   string
		char int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, test := range tests {
		first := findMarker(test.in, 14)
		if first != test.char {
			t.Errorf("Bad marker %v: %v (%v)", test.in, first, test.char)
		}
	}
}
