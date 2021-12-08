package main

import "testing"

func Test2021Day7(t *testing.T) {
	start := "16,1,2,0,4,2,7,1,2,14"

	cost := getCost(start)
	if cost != 37 {
		t.Errorf("Bad cost: %v vs 37", cost)
	}

	cost = getCostComplex(start)
	if cost != 168 {
		t.Errorf("Bad cost: %v vs 168", cost)
	}
}
