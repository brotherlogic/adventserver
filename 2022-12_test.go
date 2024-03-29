package main

import "testing"

func Test2022_12_1_Main(t *testing.T) {
	data := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	steps, route := runMap(data)
	if steps != 31 {
		t.Errorf("Bad steps: %v (31) -> %v", steps, route)
	}
}

func Test2022_12_2_Main(t *testing.T) {
	data := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	steps := runMultiMap(data)
	if steps != 29 {
		t.Errorf("Bad steps: %v (29)", steps)
	}
}
