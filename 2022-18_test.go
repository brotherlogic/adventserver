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
