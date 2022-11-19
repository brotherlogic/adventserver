package main

import "testing"

func Test2015Day20Part1(t *testing.T) {
	res := findMaxHouse(150)
	if res != 12 {
		t.Errorf("Bad max house %v (should have been 8)", res)
	}
}
