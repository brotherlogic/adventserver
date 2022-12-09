package main

import "testing"

func Test2022_9_1_Main(t *testing.T) {
	data := `R 4
	U 4
	L 3
	D 1
	R 4
	D 1
	L 5
	R 2`

	result := runRopeBridge(data)

	if result != 13 {
		t.Errorf("Bad rope bridge run: %v", result)
	}
}
