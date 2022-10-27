package main

import "testing"

func Test2015_day18_part1(t *testing.T) {
	start := `.#.#.#
	...##.
	#....#
	..#...
	#.#..#
	####..`

	now := rotate(start, 4)
	if now != 4 {
		t.Errorf("Bad rotation: %v", now)
	}
}
