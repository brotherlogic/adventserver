package main

import "testing"

func Test2015Day22Part1(t *testing.T) {
	val := magicFight(player{hitp: 10, mana: 250}, player{hitp: 13, damage: 8})
	if val != 126 {
		t.Errorf("Bad fight for num 1: %v", val)
	}

	if magicFight(player{hitp: 10, mana: 250}, player{hitp: 14, damage: 8}) != 126 {
		t.Errorf("Bad fight for num 2: %v", magicFight(player{hitp: 10, mana: 250}, player{hitp: 14, damage: 8}))
	}
}
