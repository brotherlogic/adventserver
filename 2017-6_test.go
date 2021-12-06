package main

import "testing"

func Test2017Day6Part1(t *testing.T) {
	blocks := []int{0, 2, 7, 0}

	val, val2 := computeRepeat(blocks)

	if val != 5 {
		t.Errorf("Bad repeat: %v vs 5", val)
	}

	if val2 != 4 {
		t.Errorf("Bad length: %v vs 5", val2)
	}
}
