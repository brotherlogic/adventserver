package main

import "testing"

func Test2022_8_1_Main(t *testing.T) {
	data := `30373
25512
65332
33549
35390`

	visible := countVisibleTrees(data)
	if visible != 21 {
		t.Errorf("Bad visible count: %v (21)", visible)
	}
}

func Test2022_8_2_Main(t *testing.T) {
	data := `30373
25512
65332
33549
35390`

	visible := bestTree(data)
	if visible != 8 {
		t.Errorf("Bad visible count: %v (8)", visible)
	}
}
