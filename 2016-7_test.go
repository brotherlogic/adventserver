package main

import "testing"

var cases = []struct {
	in      string
	support bool
}{
	{"abba[mnop]qrst", true},
	{"abcd[bddb]xyyx", false},
	{"aaaa[qwer]tyui", false},
	{"ioxxoj[asdfgh]zxcvbn", true},
}

func Test2016_7_1(t *testing.T) {
	for _, cas := range cases {
		support := tlsSupport(cas.in)
		if support != cas.support {
			t.Errorf("Bad support %v -> %v (was %v)", cas.in, cas.support, support)
		}
	}
}
