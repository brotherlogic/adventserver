package main

import "testing"

func Test2016_13_1(t *testing.T) {
	answer := runMaze(10, 7, 4)
	if answer != 11 {
		t.Errorf("Bad maze run: %v (11)", answer)
	}
}

func Test2016_13_2(t *testing.T) {
	answer := runMazeToLimit(4, 10)
	if answer != 9 {
		t.Errorf("Bad maze run: %v (2)", answer)
	}
}
