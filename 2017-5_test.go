package main

import "testing"

func Test2017Day5Part1(t *testing.T) {
	str := "0\n3\n0\n1\n-3"
	res := computeJumps(str)

	if res != 5 {
		t.Errorf("Bad jump %v vs 5", res)
	}
}
