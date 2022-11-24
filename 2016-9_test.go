package main

import "testing"

func Test2016_9_1(t *testing.T) {
	var cases = []struct {
		in   string
		slen int
	}{
		{"ADVENT", 6},
		{"A(1x5)BC", 7},
		{"(3x3)XYZ", 9},
		{"A(2x2)BCD(2x2)EFG", 11},
		{"(6x1)(1x3)A", 6},
		{"X(8x2)(3x3)ABCY", 18},
	}

	for _, c := range cases {
		res := expandString(c.in)
		if res != c.slen {
			t.Errorf("Bad length: %v (%v)", res, c.slen)
		}
	}
}
