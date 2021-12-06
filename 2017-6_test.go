package main

import "testing"

func Test2017Day6Part1(t *testing.T) {
	blocks := []int{0, 2, 7, 0}

	val := computeRepeat(blocks)

	if val != 5 {
		t.Errorf("Bad repeat: %v vs 5", val)
	}
}
