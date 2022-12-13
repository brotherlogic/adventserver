package main

import "testing"

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
	[1,[2,[3,[4,[5,6,0]]]],8,9]`

	sum := computeIndexSum(data)
	if sum != 13 {
		t.Errorf("Bad index sum: %v (13)", sum)
	}
}
