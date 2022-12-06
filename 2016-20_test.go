package main

import "testing"

func Test2016_20_1_Main(t *testing.T) {
	data := `5-8
0-2
4-7`
	low := getLowIp(data)

	if low != 3 {
		t.Errorf("Bad Low IP: %v (3)", low)
	}
}
