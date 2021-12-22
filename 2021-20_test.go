package main

import (
	"testing"
)

func Test2021Day20Part1(t *testing.T) {
	ieh := `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
	#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###
	.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#.
	.#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#.....
	.#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#..
	...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.....
	..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#
	`

	image := `#..#.
	#....
	##..#
	..#..
	..###`

	resolve1 := enhance(buildLarge(image), ieh)
	resolve2 := enhance(resolve1, ieh)

	printImage(resolve2)

	count := countLit(resolve2)
	if count != 35 {
		t.Errorf("Bad count: %v vs 35", count)
	}
}

func Test2021Day20Part1Focus(t *testing.T) {
	data := `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#
	
	#..#.
	#....
	##..#
	..#..
	..###`

	count := runCount(data, 2)
	if count != 35 {
		t.Errorf("Bad count: %v vs 35", count)
	}

	count2 := runCount(data, 50)
	if count2 != 3351 {
		t.Errorf("Bad count: %v vs 3351", count2)
	}
}
