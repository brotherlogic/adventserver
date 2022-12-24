package main

import "testing"

func Test2022_24_1_Main(t *testing.T) {
	data := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

	res, _ := runBlizzardMaze(buildBlizzard(data), 1)

	if res != 18 {
		t.Errorf("Bad blizzard run: %v (18)", res)
	}
}

func Test2022_24_2_Main(t *testing.T) {
	data := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

	blizz := buildBlizzard(data)
	res1, blizz1 := runBlizzardMaze(blizz, 1)
	res2, blizz2 := runBlizzardMaze(blizz1, 0)
	res3, _ := runBlizzardMaze(blizz2, 1)

	if res1+res2+res3 != 54 {
		t.Errorf("Bad blizzard run: %v (54)", res1+res2+res3)
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
