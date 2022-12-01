package main

import "testing"

func Test2022_1_1(t *testing.T) {
	data := `1000
	2000
	3000
	
	4000
	
	5000
	6000
	
	7000
	8000
	9000
	
	10000`

	max := countCalories(data)

	if max != 24000 {
		t.Errorf("Bad max: %v", max)
	}
}
