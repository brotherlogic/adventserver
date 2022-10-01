package main

import "testing"

func TestDay201512P1(t *testing.T) {
	cases := []struct {
		in     string
		skip   string
		number int
	}{
		{`[1,2,3]`, "NO", 6},
		{`{"a":2,"b":4}`, "NO", 6},
		{`[[[3]]]`, "NO", 3},
		{`{"a":{"b":4},"c":-1}`, "NO", 3},
		{`{"a":[-1,1]}`, "NO", 0},
		{`[-1,{"a":1}]`, "NO", 0},
		{`[]`, "NO", 0},
	}

	for _, c := range cases {
		count := countNumbers(c.in, c.skip)
		if count != c.number {
			t.Fatalf("Wrong number of characters on %v(%v should have been %v)", c.in, count, c.number)
		}

	}
}

func TestDay201512P2(t *testing.T) {
	cases := []struct {
		in     string
		skip   string
		number int
	}{
		{`[1,2,3]`, "red", 6},
		{`[1,{"c":"red","b":2},3]`, "red", 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, "red", 0},
		{`[1,"red",5]`, "red", 6},
	}

	for _, c := range cases {
		count := countNumbers(c.in, c.skip)
		if count != c.number {
			t.Fatalf("Wrong number of characters on %v(%v should have been %v)", c.in, count, c.number)
		}

	}
}
