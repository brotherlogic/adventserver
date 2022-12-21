package main

import (
	"log"
	"testing"
)

func Test2022_20_1_Sup(t *testing.T) {

	cases := []struct {
		in   []int
		move int
		out  []int
	}{
		{[]int{4, 5, 6, 1, 7, 8, 9}, 1, []int{4, 5, 6, 7, 1, 8, 9}},
		{[]int{4, -2, 5, 6, 7, 8, 9}, -2, []int{4, 5, 6, 7, 8, -2, 9}},
	}

	for i, cs := range cases {
		log.Printf("TEST %v", i)
		moved := moveNumber(cs.in, cs.move)

		match := true
		for i := 0; i < len(moved); i++ {
			if moved[i] != cs.out[i] {
				match = false
			}
		}

		if !match {
			t.Errorf("Bad match %v;%v -> %v (%v)", cs.in, cs.move, moved, cs.out)
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
