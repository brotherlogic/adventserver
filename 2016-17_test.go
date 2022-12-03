package main

import "testing"

func Test2016_17_1(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"ihgpwlah", "DDRRRD"},
		{"kglvqrro", "DDUDRLRRUDRD"},
		{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
	}

	for _, test := range tests {
		out := getShortestPath(test.in)
		if out != test.out {
			t.Errorf("Bad path: %v -> %v (%v)", test.in, out, test.out)
		}
	}
}
