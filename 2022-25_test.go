package main

import (
	"testing"
)

func Test2022_25_1_Single(t *testing.T) {
	val := rsnafu(1257)
	if snafu(val) != 1257 {
		t.Errorf("Bad convert: %v -> %v", 1257, val)
	}
}
func Test2022_25_1_Sup(t *testing.T) {
	cases := []struct {
		in  string
		out int64
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	}

	for _, tc := range cases {
		decimal := snafu(tc.in)
		if decimal != tc.out {
			t.Errorf("Bad convert %v -> %v (%v)", tc.in, decimal, tc.out)
		}
		sn := rsnafu(tc.out)
		if sn != tc.in {
			t.Errorf("Bad revers %v -> %v (%v)", tc.out, sn, tc.in)
		}
	}
}

func Test2022_25_1_Mainp(t *testing.T) {
	data := `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

	res := computeSnafuSum(data)
	if res != "2=-1=0" {
		t.Errorf("Bad sanfu conversion: %v (2=-1=0)", res)
	}
}
