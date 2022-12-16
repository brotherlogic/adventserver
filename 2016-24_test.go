package main

import "testing"

func Test2022_24_1_Main(t *testing.T) {
	data := `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

	steps := runDuctMaze(data)

	if steps != 14 {
		t.Errorf("Bad maze run: %v (14)", steps)
	}
}
