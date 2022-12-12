package main

import "testing"

func Test2022_12_1_Main(t *testing.T) {
	data := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	steps := runMap(data)
	if steps != 31 {
		t.Errorf("Bad steps: %v (31)", steps)
	}
}
