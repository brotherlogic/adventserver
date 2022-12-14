package main

import "testing"

func Test2022_14_1_Main(t *testing.T) {
	data := `498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9`

	sand := countSand(data)

	if sand != 24 {
		t.Errorf("Bad sand: %v (24)", sand)
	}
}

func Test2022_14_2_Main(t *testing.T) {
	data := `498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9`

	sand := countSandBased(data)

	if sand != 93 {
		t.Errorf("Bad sand: %v (93)", sand)
	}
}
