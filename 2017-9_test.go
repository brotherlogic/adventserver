package main

import "testing"

func Test2017_9_1_Main(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}}", 3},
	}

	for _, test := range tests {
		result := convertStream(test.in)
		if result != test.out {
			t.Errorf("Bad test %v -> %v (%v)", test.in, result, test.out)
		}
	}
}
