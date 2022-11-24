package main

import (
	"testing"
)

func Test2016_8_1(t *testing.T) {
	data := `rect 3x2
	rotate column x=1 by 1
	rotate row y=0 by 4
	rotate column x=1 by 1`

	result := runLightProgram(7, 3, data)

	count := 0
	for x := range result {
		for y := range result[x] {
			if result[x][y] {
				count++
			}
		}
	}

	if count != 6 {
		t.Errorf("Bad count: %v -> %v", count, result)
	}

	answer := `.#..#.#
#.#....
.#.....`

	if doPrint(result) != answer {
		t.Errorf("MISMATCH\n'%v'\n-------\n'%v'", answer, doPrint(result))
	}
}
