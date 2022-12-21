package main

import (
	"log"
	"testing"
)

func rotateMatch(a1, a2 []int) bool {
	zero1, zero2 := 0, 0

	for i := 0; i < len(a1); i++ {
		if a1[i] == 0 {
			zero1 = i
		}
		if a2[i] == 0 {
			zero2 = i
		}
	}

	for i := 0; i < len(a1); i++ {
		if a1[(zero1+i)%len(a1)] != a2[(zero2+i)%len(a1)] {
			return false
		}
	}

	return true
}

func Test2022_20_1_Sup(t *testing.T) {

	cases := []struct {
		in   []int
		move int
		out  []int
	}{
		{[]int{4, 5, 6, 1, 7, 8, 9}, 1, []int{4, 5, 6, 7, 1, 8, 9}},
		{[]int{4, -2, 5, 6, 7, 8, 9}, -2, []int{4, 5, 6, 7, 8, -2, 9}},
		{[]int{1, 2, -3, 3, -2, 0, 4}, 1, []int{2, 1, -3, 3, -2, 0, 4}},
		{[]int{2, 1, -3, 3, -2, 0, 4}, 2, []int{1, -3, 2, 3, -2, 0, 4}},
		{[]int{1, 2, -2, -3, 0, 3, 4}, -2, []int{1, 2, -3, 0, 3, 4, -2}},
		{[]int{1, 2, -3, 0, 3, 4, -2}, 4, []int{1, 2, -3, 4, 0, 3, -2}},
	}

	for i, cs := range cases {
		log.Printf("RUN %v", cs.in)
		moved := moveNumber(cs.in, cs.move)

		match := rotateMatch(moved, cs.out)

		if !match {
			t.Errorf("Bad match (%v) %v;%v -> %v (%v)", i, cs.in, cs.move, moved, cs.out)
		}
	}
}

func Test2022_20_1_Steps(t *testing.T) {

	cases := []struct {
		in   []int
		move int
		out  []int
	}{
		{[]int{1, 2, -3, 3, -2, 0, 4}, 1, []int{2, 1, -3, 3, -2, 0, 4}},
		{[]int{2, 1, -3, 3, -2, 0, 4}, 2, []int{1, -3, 2, 3, -2, 0, 4}},
		{[]int{1, -3, 2, 3, -2, 0, 4}, -3, []int{1, 2, 3, -2, -3, 0, 4}},
		{[]int{1, 2, 3, -2, -3, 0, 4}, 3, []int{1, 2, -2, -3, 0, 3, 4}},
		{[]int{1, 2, -2, -3, 0, 3, 4}, -2, []int{1, 2, -3, 0, 3, 4, -2}},
		{[]int{1, 2, -3, 0, 3, 4, -2}, 0, []int{1, 2, -3, 0, 3, 4, -2}},
		{[]int{1, 2, -3, 0, 3, 4, -2}, 4, []int{1, 2, -3, 4, 0, 3, -2}},
	}

	for i, cs := range cases {
		log.Printf("RUN %v", cs.in)
		moved := moveNumber(cs.in, cs.move)

		match := rotateMatch(moved, cs.out)

		if !match {
			t.Errorf("Bad match (%v) %v;%v -> %v (%v)", i, cs.in, cs.move, moved, cs.out)
		}
	}
}

func Test2022_20_1_Main(t *testing.T) {
	data := `1
	2
	-3
	3
	-2
	0
	4`

	num := unencrpyt(data)
	if num != 3 {
		t.Errorf("Bad unencrpyption: %v (3)", num)
	}
}
