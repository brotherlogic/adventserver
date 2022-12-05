package main

import "testing"

func Test2022_5_1(t *testing.T) {
	data := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
	
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	re := rearrangeCrates(data)

	if re != "CMZ" {
		t.Errorf("Bad rearrange: %v (CMZ)", re)
	}
}
