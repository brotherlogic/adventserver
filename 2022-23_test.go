package main

import (
	"math"
	"testing"
)

func Test2022_23_2_Main(t *testing.T) {
	data := `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

	result := runElves(data, math.MaxInt)
	if result != 20 {
		t.Errorf("Bad times: %v (20)", result)
	}
}

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

func Test2022_23_1_Sup(t *testing.T) {
	data := `.....
..##.
..#..
.....
..##.
.....`

	result := runElves(data, 3)
	if result != 30-5 {
		t.Errorf("Bad times: %v (25)", result)
	}
}
