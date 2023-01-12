package main

import "testing"

func Test2017_12_1_Main(t *testing.T) {

	data := `0 <-> 2
	1 <-> 1
	2 <-> 0, 3, 4
	3 <-> 2, 4
	4 <-> 2, 3, 6
	5 <-> 6
	6 <-> 4, 5`

	res := countZeros(data)
	if res != 6 {
		t.Errorf("Wrong zero count %v (should be 6)", res)
	}
}
