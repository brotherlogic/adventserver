package main

import "testing"

func TestDay9of2021(t *testing.T) {
	data := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	compute := getRisk(data)

	if compute != 15 {
		t.Errorf("Compute is wrong %v vs 15", compute)
	}
}
