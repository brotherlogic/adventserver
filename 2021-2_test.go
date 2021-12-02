package main

import "testing"

func TestDay2Part1(t *testing.T) {
	data := "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"

	if procSub(data) != 150 {
		t.Errorf("Bad result -> %v vs %v", procSub(data), 150)
	}
}
