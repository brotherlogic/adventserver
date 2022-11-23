package main

import "testing"

func Test2015_24_1(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	best := computeGrouping(weights)

	if best != 99 {
		t.Errorf("Bad grouping: %v (99)", best)
	}
}

func TestSup2015_24_1(t *testing.T) {
	weights := []int{1, 2, 3, 5, 7, 13, 17, 19, 23, 29, 31, 37, 41, 43, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
	best := computeGrouping(weights)

	if best != 99 {
		t.Errorf("Bad grouping: %v (99)", best)
	}
}
