package main

import "testing"

func Test2015Day19Part2(t *testing.T) {
	data := `e => H
	e => O
	H => HO
	H => OH
	O => HH
	
	HOH`

	steps := treeMolecules(data)

	if steps != 3 {
		t.Errorf("Bad Tree: %v (Should be 3)", steps)
	}
}

func TestSup2015Day19Part2(t *testing.T) {
	data := `e => H
	e => O
	H => HO
	H => OH
	O => HH
	
	HOHOHO`

	steps := treeMolecules(data)

	if steps != 6 {
		t.Errorf("Bad Tree: %v (Should be 6)", steps)
	}
}

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
