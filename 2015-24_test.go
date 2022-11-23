package main

import (
	"testing"
)

func Test2015_24_1(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	best := computeGrouping(weights, 3)

	if best != 99 {
		t.Errorf("Bad grouping: %v (99)", best)
	}
}
func Test2015_24_2(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	best := computeGrouping(weights, 4)

	if best != 44 {
		t.Errorf("Bad grouping: %v (44)", best)
	}
}
