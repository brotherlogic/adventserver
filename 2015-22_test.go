package main

import "testing"

func Test2015Day22Part1(t *testing.T) {
	if !magicFight(player{hitp: 10, mana: 250}, player{hitp: 13, damage: 8}) {
		t.Errorf("Bad fight for num 1")
	}

	if !magicFight(player{hitp: 10, mana: 250}, player{hitp: 14, damage: 8}) {
		t.Errorf("Bad fight for ")
	}
}
