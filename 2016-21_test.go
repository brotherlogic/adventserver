package main

import "testing"

func Test2016_21_1(t *testing.T) {
	tests := []struct {
		in   string
		code string
		out  string
	}{
		{"abcde", "swap position 4 with position 0", "ebcda"},
		{"ebcda", "swap letter d with letter b", "edcba"},
		{"edcba", "reverse positions 0 through 4", "abcde"},
		{"abcde", "rotate left 1 step", "bcdea"},
		{"bcdea", "move position 1 to position 4", "bdeac"},
		{"bdeac", "move position 3 to position 0", "abdec"},
		{"abdec", "rotate based on position of letter b", "ecabd"},
		{"ecabd", "rotate based on position of letter d", "decab"},
	}

	for _, test := range tests {
		out := translateCode(test.in, test.code)
		if out != test.out {
			t.Errorf("Bad translate %v -> %v => %v (%v)", test.in, test.code, out, test.out)
		}
	}
}

func Test2016_21_2(t *testing.T) {
	data := `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`

	result := "decab"

	start := generateAll("abcde", "", result, data)

	if start != "abcde" {
		t.Errorf("Bad result: %v", start)
	}
}
