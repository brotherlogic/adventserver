package main

import "testing"

func Test2016_11_1(t *testing.T) {
	data := `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
	The second floor contains a hydrogen generator.
	The third floor contains a lithium generator.
	The fourth floor contains nothing relevant.`

	res, _ := findFloors(data)
	if res != 11 {
		t.Errorf("Could not resolve program: %v (%v)", res, 11)
	}
}
