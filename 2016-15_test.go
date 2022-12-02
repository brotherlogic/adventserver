package main

import "testing"

func Test2016_15_1(t *testing.T) {
	data := `Disc #1 has 5 positions; at time=0, it is at position 4.
	Disc #2 has 2 positions; at time=0, it is at position 1.`

	result := fallBalls(data, false)

	if result != 5 {
		t.Errorf("Bad resolution: %v", result)
	}
}
