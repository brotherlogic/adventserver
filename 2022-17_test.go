package main

import (
	"log"
	"testing"
)

func Test2022_17_1_Main(t *testing.T) {
	data := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	tetis := 2022
	r, _ := runTetris(data, tetis)
	res := getHeight(r)

	if res != 3068 {
		t.Errorf("Bad height: %v (3068): %v", res, getHeight(r))
	}
}

func Test2022_17_2_Basic(t *testing.T) {
	data := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	tetis := 10000
	r, chamber := runTetris(data, tetis)
	res := getHeight(r)

	rep := findRepeat(chamber, res-20, 2022)
	log.Printf("LEN: %v", len(data))

	if rep != 3069 {
		t.Errorf("Bad height: %v (3068): %v", rep, getHeight(r))
	}
}

func Test2022_17_2_Main(t *testing.T) {
	data := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	tetis := 200
	r, chamber := runTetris(data, tetis)
	res := getHeight(r)

	rep := findRepeat(chamber, res-20, 1000000000000)
	log.Printf("LEN: %v", len(data))

	if rep != 1514285714288 {
		t.Errorf("Bad height: %v (1514285714288): %v", rep, getHeight(r))
	}
}
