package main

import (
	"fmt"
	"testing"
)

func Test2017_10_1_Sup(t *testing.T) {
	tests := []struct {
		in             []int
		range1, range2 int
		out            []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 5, []int{5, 4, 3, 2, 1, 0, 6, 7, 8, 9, 10}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 4, []int{4, 3, 2, 1, 0, 5, 6, 7, 8, 9, 10}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 10, 9, 8}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 0, []int{8, 1, 2, 3, 4, 5, 6, 7, 0, 10, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 1, []int{9, 8, 2, 3, 4, 5, 6, 7, 1, 0, 10}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 1, []int{10, 9, 2, 3, 4, 5, 6, 7, 8, 1, 0}},
	}

	for _, test := range tests {
		rev := reverseArr(test.in, test.range1, test.range2)
		if fmt.Sprintf("%v", test.out) != fmt.Sprintf("%v", rev) {
			t.Errorf("%v (%v,%v) -> %v [%v]", test.in, test.range1, test.range2, rev, test.out)
		}
	}
}

func Test2017_10_1_Main(t *testing.T) {
	res := runArray([]int{0, 1, 2, 3, 4}, []int{3, 4, 1, 5, 0})

	if res != 12 {
		t.Errorf("Bad Array run: %v (12)", res)
	}
}
