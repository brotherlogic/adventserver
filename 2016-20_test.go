package main

import (
	"context"
	"fmt"
	"testing"
)

func Test2016_20_1_Main(t *testing.T) {
	data := `5-8
0-2
4-7`
	low := getLowIp(data)

	if low != 3 {
		t.Errorf("Bad Low IP: %v (3)", low)
	}
}

func Test2016_20_2_Main(t *testing.T) {
	data := `5-8
0-2
4-7`
	low := getIps(context.Background(), data, 9, tlog)

	if low != "2" {
		t.Errorf("Bad Low IP: %v (2)", low)
	}
}

func Test2016_20_2_Sup(t *testing.T) {
	data := `5-8
0-2
4-7`
	low := getIps(context.Background(), data, 4294967295, tlog)

	if low != fmt.Sprintf("%v", 4294967295-9+2) {
		t.Errorf("Bad Low IP: %v (%v)", low, 4294967295-9+2)
	}
}

func Test2016_20_2_Range(t *testing.T) {
	data := `3872622758-3876157728
4222455331-4230019102`
	low := getIps(context.Background(), data, 4294967295, tlog)

	if low != "411245795" {
		t.Errorf("Bad Low IP: %v (%v)", low, 411245795)
	}
}

func Test2016_20_1_Sup(t *testing.T) {
	data := `5-8
0-2
3-7`
	low := getLowIp(data)

	if low != 9 {
		t.Errorf("Bad Low on IP: %v (9)", low)
	}
}

func Test2016_20_1_Sup2(t *testing.T) {
	data := `2179314-4534265
2171134-4793563
0-2179314`
	low := getLowIp(data)

	if low != 4793564 {
		t.Errorf("Bad Low IP on biggy: %v (4793564)", low)
	}
}

func Test2016_20_1_Inside(t *testing.T) {
	data := `0-5
2-3`
	low := getLowIp(data)

	if low != 6 {
		t.Errorf("Bad Low IP on biggy: %v (6)", low)
	}
}

func Test2016_20_1_Outside(t *testing.T) {
	data := `2-3
0-5`
	low := getLowIp(data)

	if low != 6 {
		t.Errorf("Bad Low IP on biggy: %v (6)", low)
	}
}

func Test2016_20_1_OverLeft(t *testing.T) {
	data := `0-3
2-5`
	low := getLowIp(data)

	if low != 6 {
		t.Errorf("Bad Low IP on biggy: %v (6)", low)
	}
}

func Test2016_20_1_OverRight(t *testing.T) {
	data := `2-5
0-3`
	low := getLowIp(data)

	if low != 6 {
		t.Errorf("Bad Low IP on biggy: %v (6)", low)
	}
}

func Test2016_20_1_TightLeft(t *testing.T) {
	data := `0-3
3-5`
	low := getLowIp(data)

	if low != 6 {
		t.Errorf("Bad Low IP on biggy: %v (6)", low)
	}
}

func Test2016_20_1_TightRight(t *testing.T) {
	data := `3-5
0-3`
	low := getLowIp(data)

	if low != 6 {
		t.Errorf("Bad Low IP on biggy: %v (6)", low)
	}
}
