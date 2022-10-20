package main

import "testing"

func Test2015_16_1(t *testing.T) {
	known := properties{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	aunt := findAunt("Sue 1: cars: 9, akitas: 3, goldfish: 0", known)

	if aunt != 0 {
		t.Errorf("Found aunt?")
	}

	aunt = findAunt("Sue 1: cars: 2, akitas: 0, goldfish: 5", known)

	if aunt != 1 {
		t.Errorf("Found aunt: %v", aunt)
	}
}
