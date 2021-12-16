package main

import "testing"

func Test2021Day15(t *testing.T) {
	data := `1163751742
	1381373672
	2136511328
	3694931569
	7463417111
	1319128137
	1359912421
	3125421639
	1293138521
	2311944581`

	path := getBestPath(data)
	if path != 40 {
		t.Errorf("Bad Path %v vs 40", path)
	}
}
