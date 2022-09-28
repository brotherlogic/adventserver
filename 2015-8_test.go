package main

import "testing"

func TestDay20158P1(t *testing.T) {
	cases := []struct {
		in         string
		characters int
		codes      int
	}{
		{`""`, 2, 0},
		{`"abc"`, 5, 3},
		{`"aaa\"aaa"`, 10, 7},
		{`"\x27"`, 6, 1},
		{`"aaa\"aaa\x27"`, 14, 8},
		{`"\\\"`, 6, 2},
	}

	for _, c := range cases {
		chs, cds := computeStringLength(c.in)
		if chs != c.characters {
			t.Fatalf("Wrong number of characters on %v(%v should have been %v)", c.in, chs, c.characters)
		}
		if cds != c.codes {
			t.Fatalf("Wrong string length on %v (%v should have been %v)", c.in, cds, c.codes)
		}
	}
}
