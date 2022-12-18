package main

import "testing"

func Test2022_17_1_Main(t *testing.T) {
	data := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	res := getHeight(runTetris(data))

	if res != 3068 {
		t.Errorf("Bad height: %v (3068)", res)
	}
}
