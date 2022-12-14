package main

import (
	"context"
	"testing"
)

func Test2022_13_1_Main(t *testing.T) {
	data := `[1,1,3,1,1]
	[1,1,5,1,1]
	
	[[1],[2,3,4]]
	[[1],4]
	
	[9]
	[[8,7,6]]
	
	[[4,4],4,4]
	[[4,4],4,4,4]
	
	[7,7,7,7]
	[7,7,7]
	
	[]
	[3]
	
	[[[]]]
	[[]]
	
	[1,[2,[3,[4,[5,6,7]]]],8,9]
	[1,[2,[3,[4,[5,6,0]]]],8,9]
	
	[10,2]
	[10,1]

	[[10,6,9,4,[]]]
	[[[],[8,[0,9]],7],[5,1],[5,[],[10,9],0,8]]`

	sum := computeIndexSum(context.Background(), data, tlog)
	if sum != 13 {
		t.Errorf("Bad index sum: %v (13)", sum)
	}
}

func Test2022_13_1_Focus(t *testing.T) {
	data := `[[],[1,[3,[0]]],[]]
	[[],[[0,[5,3,0,1,0],[3,0,5,7],10,[2,8,5,0]],10,[2,4,[1],[5,6,7],[]],[]]]`

	sum := computeIndexSum(context.Background(), data, tlog)
	if sum != 0 {
		t.Errorf("Bad index sum: %v (0)", sum)
	}
}

func Test2022_13_2_Main(t *testing.T) {
	data := `[1,1,3,1,1]
	[1,1,5,1,1]
	
	[[1],[2,3,4]]
	[[1],4]
	
	[9]
	[[8,7,6]]
	
	[[4,4],4,4]
	[[4,4],4,4,4]
	
	[7,7,7,7]
	[7,7,7]
	
	[]
	[3]
	
	[[[]]]
	[[]]
	
	[1,[2,[3,[4,[5,6,7]]]],8,9]
	[1,[2,[3,[4,[5,6,0]]]],8,9]
	
	[10,2]
	[10,1]`

	sum := resolvePackets(data)
	if sum != 140 {
		t.Errorf("Bad index sum: %v (140)", sum)
	}
}
