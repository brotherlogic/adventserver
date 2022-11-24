package main

import "testing"

func Test2016_7_1(t *testing.T) {
	var cases = []struct {
		in      string
		support bool
	}{
		{"abba[mnop]qrst", true},
		{"abcd[bddb]xyyx", false},
		{"aaaa[qwer]tyui", false},
		{"ioxxoj[asdfgh]zxcvbn", true},
	}

	for _, cas := range cases {
		support := tlsSupport(cas.in)
		if support != cas.support {
			t.Errorf("Bad support %v -> %v (was %v)", cas.in, cas.support, support)
		}
	}
}

func Test2016_7_2(t *testing.T) {
	var cases = []struct {
		in      string
		support bool
	}{
		{"aba[bab]xyz", true},
		{"xyx[xyx]xyx", false},
		{"aaa[kek]eke", true},
		{"zazbz[bzb]cdb", true},
	}

	for _, cas := range cases {
		support := sslSupport(cas.in)
		if support != cas.support {
			t.Errorf("Bad support %v -> %v (was %v)", cas.in, cas.support, support)
		}
	}
}
