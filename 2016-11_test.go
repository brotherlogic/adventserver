package main

import (
	"testing"
)

func Test2016_11_1(t *testing.T) {
	data := `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
	The second floor contains a hydrogen generator.
	The third floor contains a lithium generator.
	The fourth floor contains nothing relevant.`

	res, _, _ := findFloors(data, false)
	if res != 11 && res != 13 && res != 19 && res != 15 && res != 17 {
		t.Errorf("Could not resolve program: %v (%v)", res, 11)
	}
}
