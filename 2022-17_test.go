package main

import "testing"

func Test2022_17_1_Main(t *testing.T) {
	data := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	tetis := 2022
	res := getHeight(runTetris(data, tetis))

	if res != 3068 {
		t.Errorf("Bad height: %v (3068): %v", res, runTetris(data, tetis))
	}
}
