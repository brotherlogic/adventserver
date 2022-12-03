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

func Test2016_17_2(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"ihgpwlah", 370},
		{"kglvqrro", 492},
		{"ulqzkmiv", 830},
	}

	for _, test := range tests {
		out := getLongestPath(test.in)
		if out != test.out {
			t.Errorf("Bad path: %v -> %v (%v)", test.in, out, test.out)
		}
	}
}
