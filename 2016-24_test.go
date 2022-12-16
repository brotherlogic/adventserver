package main

import "testing"

func Test2016_24_1_Main(t *testing.T) {
	data := `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

	steps := runDuctMaze(data, false)

	if steps != 14 {
		t.Errorf("Bad maze run: %v (14)", steps)
	}
}

func Test2016_24_2_Main(t *testing.T) {
	data := `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

	steps := runDuctMaze(data, true)

	if steps != 20 {
		t.Errorf("Bad maze run: %v (20)", steps)
	}
}
