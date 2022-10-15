package main

import "testing"

func Test201515Part1(t *testing.T) {
	data := `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

	bestScore := computeBestScore(data)

	if bestScore != 62842880 {
		t.Errorf("Score should have been 62842880, was %v", bestScore)
	}
}
