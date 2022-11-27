package main

import "testing"

func Test2016_11_1(t *testing.T) {
	data := `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
	The second floor contains a hydrogen generator.
	The third floor contains a lithium generator.
	The fourth floor contains nothing relevant.`

	res, _, _ := findFloors(data)
	if res != 11 {
		t.Errorf("Could not resolve program: %v (%v)", res, 11)
	}
}

func Test2016_11_1_read(t *testing.T) {
	data := `The first floor contains a strontium generator, a strontium-compatible microchip, a plutonium generator, and a plutonium-compatible microchip.
	The second floor contains a thulium generator, a ruthenium generator, a ruthenium-compatible microchip, a curium generator, and a curium-compatible microchip.
	The third floor contains a thulium-compatible microchip.
	The fourth floor contains nothing relevant.`

	found := false
	floors := buildFloors(data)
	for _, f := range floors.floors {
		for _, e := range f {
			if e == "rM" {
				found = true
			}
		}
	}

	if !found {
		t.Errorf("Did not find rM: %+v", floors)
	}
}
