package main

import "testing"

func Test2022_18_1_Basic(t *testing.T) {
	data := `1,1,1
	2,1,1`

	edges := countEdges(data)
	if edges != 10 {
		t.Errorf("Bad edges: %v (10)", edges)
	}
}

func Test2022_18_1_Main(t *testing.T) {
	data := `2,2,2
	1,2,2
	3,2,2
	2,1,2
	2,3,2
	2,2,1
	2,2,3
	2,2,4
	2,2,6
	1,2,5
	3,2,5
	2,1,5
	2,3,5`

	edges := countEdges(data)
	if edges != 64 {
		t.Errorf("Bad edges: %v (64)", edges)
	}
}

func Test2022_18_2_Main(t *testing.T) {
	data := `2,2,2
	1,2,2
	3,2,2
	2,1,2
	2,3,2
	2,2,1
	2,2,3
	2,2,4
	2,2,6
	1,2,5
	3,2,5
	2,1,5
	2,3,5`

	edges := countEdgesExt(data)
	if edges != 58 {
		t.Errorf("Bad edges: %v (58)", edges)
	}
}
