package main

import "testing"

func Test2021Day17(t *testing.T) {
	bestx, besty, height := findBest(20, 30, -10, -5)

	if height != 45 {
		t.Errorf("Bad result %v, %v, %v", bestx, besty, height)
	}
}

func Test2021Day17Part1Focus(t *testing.T) {
	cases := []struct {
		x      int
		y      int
		height int
	}{
		{7, 2, 3},
		{6, 3, 6},
		{9, 0, 0},
		{17, -4, 0},
		{6, 9, 45},
	}

	for _, cs := range cases {
		height := throwIn(cs.x, cs.y, 20, -10, 30, -5)
		if height != cs.height {
			t.Errorf("Error in %v,%v -> %v (%v)", cs.x, cs.y, height, cs.height)
		}
	}
}
