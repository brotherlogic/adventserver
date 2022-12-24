package main

import "testing"

func Test2022_24_1_Main(t *testing.T) {
	data := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

	res := runBlizzardMaze(data)

	if res != 18 {
		t.Errorf("Bad blizzard run: %v (18)", res)
	}
}

func Test2022_24_1_Sup(t *testing.T) {
	data := `#.#####
	#.....#
	#.>...#
	#.....#
	#.....#
	#...v.#
	#####.#`

	blizz := buildBlizzard(data)
	next := blizz.next()

	if next.bx[0] != 3 && next.by[0] != 2 {
		t.Errorf("Bad cycle: %+v", next)
	}
}
