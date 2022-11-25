package main

import (
	"testing"
)

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
			t.Errorf("Bad length #1: %v (%v)", res, c.slen)
		}
	}
}

func Test2016_9_2(t *testing.T) {
	var cases = []struct {
		in   string
		slen int64
	}{
		{"(3x3)XYZ", 9},
		{"X(8x2)(3x3)ABCY", int64(len("XABCABCABCABCABCABCY"))},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for _, c := range cases {
		res := searchString(c.in)
		if res != c.slen {
			t.Errorf("Bad length #2: %v (%v)", res, c.slen)
		}
	}
}
