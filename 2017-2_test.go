package main

import "testing"

func Test2017Day2Part1(t *testing.T) {
	input := "5 1 9 5\n7 5 3\n2 4 6 8"

	res := getChecksum(input)

	if res != int32(18) {
		t.Errorf("Bad checksum: %v vs %v", res, 18)
	}
}
