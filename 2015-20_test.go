package main

import "testing"

func Test2015Day20Part1(t *testing.T) {
	res := findMaxHouse(3600000)
	if res != 90720 {
		t.Errorf("Bad max house %v (should have been 8)", res)
	}
}
