package main

import "testing"

func Test2022_2_1(t *testing.T) {
	data := `A Y
	B X
	C Z`

	score := getRPSScore(data)

	if score != 15 {
		t.Errorf("Bad RPS Score: %v (15)", score)
	}
}
