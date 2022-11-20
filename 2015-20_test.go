package main

import "testing"

func Test2015Day20Part1(t *testing.T) {
	res := findMaxHouse(15)
	if res != 8 {
		t.Errorf("Bad max house %v (should have been 8)", res)
	}
}
