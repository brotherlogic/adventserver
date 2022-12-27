package main

import "testing"

func Test2017_8_1_Main(t *testing.T) {
	data := `b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10`

	res := runJumpProgram(data)

	highest := 0
	for _, value := range res.register {
		if value > highest {
			highest = value
		}
	}

	if highest != 1 {
		t.Errorf("Bad registeR: %v", highest)
	}
}

func Test2017_8_2_Main(t *testing.T) {
	data := `b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10`

	res := runJumpProgram(data)

	highest := 0
	for _, value := range res.hregister {
		if value > highest {
			highest = value
		}
	}

	if highest != 10 {
		t.Errorf("Bad registeR: %v", highest)
	}
}
