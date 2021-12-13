package main

import (
	"fmt"
	"testing"
)

var (
	data = `6,10
	0,14
	9,10
	0,3
	10,4
	4,11
	6,0
	6,12
	4,1
	0,13
	10,12
	3,4
	3,0
	8,4
	1,10
	2,14
	8,10
	9,0
	
	fold along y=7
	fold along x=5`
)

func Test2021Day13(t *testing.T) {
	val := runFolds(data, 1)
	if val != 17 {
		t.Errorf("Wrong dots %v vs 17", val)
	}
}

func Test2021Day13Specific(t *testing.T) {
	grid := buildGrid(data)
	printGrid(grid)
	fmt.Printf("---\n")
	grid = horizFold(grid, 7)
	printGrid(grid)
	fmt.Printf("---\n")
	grid = vertFold(grid, 5)
	printGrid(grid)
}
