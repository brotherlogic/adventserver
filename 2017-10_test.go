package main

import "testing"

func Test2017_10_1_Main(t *testing.T) {
	res := runArray([]int{0, 1, 2, 3, 4}, []int{3, 4, 1, 5})

	if res != 12 {
		t.Errorf("Bad Array run: %v (12)", res)
	}
}
