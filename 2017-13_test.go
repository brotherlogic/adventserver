package main

import "testing"

func Test2017_13_1_Main(t *testing.T) {
	data := `0: 3
	1: 2
	4: 4
	6: 4`

	res := computeSeverity(data)
	if res != 24 {
		t.Errorf("Bad severity (should have been 24): %v", res)
	}
}
