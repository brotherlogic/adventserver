package main

import "testing"

func Test2015Day19Part1(t *testing.T) {
	data := `H => HO
	H => OH
	O => HH
	
	HOH`

	num := getMolecules(data)

	if num != 4 {
		t.Errorf("Bad moles: %v -> 4", num)
	}
}

func Test2015Day19Part1Other(t *testing.T) {
	data := `H => HO
	H => OH
	O => HH
	
	HOHOHO`

	num := getMolecules(data)

	if num != 7 {
		t.Errorf("Bad moles: %v -> 7", num)
	}
}
