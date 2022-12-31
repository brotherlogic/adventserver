package main

import (
	"fmt"
	"strconv"
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

func Test2017_10_2_Sup(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for _, test := range tests {
		rev := runHash(test.in, []int{17, 31, 73, 47, 23})
		if rev != test.out {
			t.Errorf("Bad run: %v -> [%v] (%v)", test.in, rev, test.out)
		}
	}
}
func Test2017_10_2_Sparse(t *testing.T) {
	arr := []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}
	val := arr[0]
	for i := 1; i < len(arr); i++ {
		val ^= arr[i]
	}

	if val != 64 {
		t.Errorf("Bad XOR: %v (64)", val)
	}

	narr := []int{64, 7, 255}
	ret := ""
	for _, val := range narr {
		v := strconv.FormatInt(int64(val), 16)
		if len(v) == 1 {
			ret += "0" + v
		} else {
			ret += v
		}
	}

	if ret != "4007ff" {
		t.Errorf("Bad hex: %v (4007ff)", ret)
	}
}
