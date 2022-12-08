package main

import "testing"

func Test2016_22_1(t *testing.T) {
	data := "/dev/grid/node-x0-y0     91T   66T    25T   72%"

	nodes := buildNodes(data)

	if nodes[0].used != 66 {
		t.Errorf("Bad node: %+v", nodes[0])
	}
}
