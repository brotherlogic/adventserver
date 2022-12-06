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

func Test2016_20_1_Sup(t *testing.T) {
	data := `5-8
0-2
3-7`
	low := getLowIp(data)

	if low != 9 {
		t.Errorf("Bad Low IP: %v (9)", low)
	}
}

func Test2016_20_1_Sup2(t *testing.T) {
	data := `2179314-4534265
2171134-4793563
0-2179314`
	low := getLowIp(data)

	if low != 9 {
		t.Errorf("Bad Low IP: %v (9)", low)
	}
}
