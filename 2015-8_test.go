package main

import "testing"

func TestDay20158P1(t *testing.T) {
	cases := []struct {
		in         string
		characters int
		codes      int
	}{
		{`""`, 2, 0},
		{`"abc"`, 5, 0},
		{`"aaa\"aaa"`, 10, 7},
		{`"\x27"`, 6, 1},
	}

	for _, c := range cases {
		chs, cds := computeStringLength(c.in)
		if chs != c.characters {
			t.Fatalf("Wrong number of characters (%v should have been %v)", chs, c.characters)
		}
		if cds != c.codes {
			t.Fatalf("Wrong string length (%v should have been %v)", cds, c.codes)
		}
	}
}
