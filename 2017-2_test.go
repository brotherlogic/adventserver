package main

import "testing"

func Test2017Day2Part1(t *testing.T) {
	input := "5 1 9 5\n7 5 3\n2 4 6 8"

	res := getChecksum(input)

	if res != int32(18) {
		t.Errorf("Bad checksum: %v vs %v", res, 18)
	}
}

func Test2017Day2Part2(t *testing.T) {
	input := "5 9 2 8\n9 4 7 3\n3 8 6 5"

	res := evenDivide(input)

	if res != int32(9) {
		t.Errorf("Bad checksum: %v vs %v", res, 9)
	}
}
