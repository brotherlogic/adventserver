package main

import (
	"testing"
)

func Test2021Day11(t *testing.T) {
	data := `11111
	19991
	19191
	19991
	11111`

	arr := buildArr(data)
	count := flash(arr)
	if count != 9 {
		t.Errorf("Did not flash the riight number: %v and \n%v", count, restoreString(arr))
	}
}

func Test2021Specific(t *testing.T) {

	data2 := `8807476555
	5089087054
	8597889608
	8485769600
	8700908800
	6600088989
	6800005943
	0000007456
	9000000876
	8700006848`

	arr2 := buildArr(data2)
	flash(arr2)
}

func Test2021Day11Part2(t *testing.T) {
	data := `5483143223
	2745854711
	5264556173
	6141336146
	6357385478
	4167524645
	2176841721
	6882881134
	4846848554
	5283751526`

	arr := buildArr(data)
	count := 0
	for i := 0; i < 10; i++ {
		count += flash(arr)
	}

	if count != 204 {
		t.Errorf("Bad flash count: %v vs 204", count)
	}

	for i := 0; i < 90; i++ {
		count += flash(arr)
	}

	if count != 1656 {
		t.Errorf("Bad flash count: %v vs 1656", count)
	}
}

func Test2021Day11ActualPart2(t *testing.T) {
	data := `5483143223
	2745854711
	5264556173
	6141336146
	6357385478
	4167524645
	2176841721
	6882881134
	4846848554
	5283751526`

	arr := buildArr(data)

	seenAt := -1
	for i := 1; i <= 1000; i++ {
		if seenAt < 0 && flash(arr) == 100 {
			seenAt = i
		}
	}

	if seenAt != 195 {
		t.Errorf("Seen wrong %v vs 195", seenAt)
	}
}
