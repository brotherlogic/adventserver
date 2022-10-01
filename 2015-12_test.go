package main

import "testing"

func TestDay201512P1(t *testing.T) {
	cases := []struct {
		in     string
		number int
	}{
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
	}

	for _, c := range cases {
		count := countNumbers(c.in)
		if count != c.number {
			t.Fatalf("Wrong number of characters on %v(%v should have been %v)", c.in, count, c.number)
		}

	}
}
