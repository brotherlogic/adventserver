package main

import "testing"

func Test2021Day21Part1(t *testing.T) {
	loser, rolls := runGame(4, 8, 1000)
	if rolls*loser != 739785 {
		t.Errorf("Bad game %v and %v", loser, rolls)
	}
}
