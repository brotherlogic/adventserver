package main

import "testing"

func Test2017_11_1_Main(t *testing.T) {

	cases := []struct {
		in  string
		out int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}

	for _, test := range cases {
		steps := computeSteps(test.in)
		if steps != test.out {
			t.Errorf("Bad steps: %v -> %v (%v)", test.in, steps, test.out)
		}
	}
}
