package main

import "testing"

func Test2021Day7Part1(t *testing.T) {
	days1 := computeCycle("3,4,3,1,2", 18)
	if days1 != 26 {
		t.Errorf("Bad count: %v vs 26", days1)
	}

	days2 := computeCycle("3,4,3,1,2", 80)
	if days2 != 5934 {
		t.Errorf("Bad count: %v vs 5934", days2)
	}

	days3 := computeCycle("3,4,3,1,2", 256)
	if days3 != 26984457539 {
		t.Errorf("Bad count: %v vs 26984457539", days2)
	}
}
