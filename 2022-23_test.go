package main

import "testing"

func Test2022_23_1_Main(t *testing.T) {
	data := `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

	result := runElves(data, 10)
	if result != 110 {
		t.Errorf("Bad times: %v (110)", result)
	}
}
