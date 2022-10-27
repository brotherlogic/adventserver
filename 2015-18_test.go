package main

import "testing"

func Test2015_day18_part1(t *testing.T) {
	start := `.#.#.#
	...##.
	#....#
	..#...
	#.#..#
	####..`

	now := rotate(start, 4, false)
	if now != 4 {
		t.Errorf("Bad rotation: %v", now)
	}
}

func Test2015_day18_part2(t *testing.T) {
	start := `.#.#.#
	...##.
	#....#
	..#...
	#.#..#
	####..`

	now := rotate(start, 5, true)
	if now != 17 {
		t.Errorf("Bad rotation: %v", now)
	}
}
