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
