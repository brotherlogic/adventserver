package main

import "testing"

func TestPart1(t *testing.T) {
	p1 := player{hitp: 8, damage: 5, armor: 5}
	p2 := player{hitp: 12, damage: 7, armor: 2}

	if !fight(p1, p2) {
		t.Errorf("Bad battle")
	}
}
