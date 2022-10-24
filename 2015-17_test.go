package main

import "testing"

func Test2015_Day17_Part1(t *testing.T) {
	containers := `20
15
10
5
5`

	val := computeContainers(containers, 25)
	if val != 4 {
		t.Errorf("Bad containers: %v -> 4", val)
	}
}
