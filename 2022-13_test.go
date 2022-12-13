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
	[10,1]`

	sum := computeIndexSum(context.Background(), data, tlog)
	if sum != 13 {
		t.Errorf("Bad index sum: %v (13)", sum)
	}
}
