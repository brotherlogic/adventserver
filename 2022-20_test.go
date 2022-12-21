package main

import "testing"

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
