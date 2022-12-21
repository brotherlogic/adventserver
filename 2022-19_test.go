package main

import "testing"

func Test2022_19_1_Main(t *testing.T) {
	data := `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
  Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

	bestBlue := getBestBlue(data)

	if bestBlue != 33 {
		t.Errorf("Bad blueprints: %v (33)", bestBlue)
	}
}

func Test2022_19_1_Sup1(t *testing.T) {
	data := `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.`

	bestBlue := getBestBlue(data)

	if bestBlue != 9 {
		t.Errorf("Bad blueprints: %v (33)", bestBlue)
	}
}

func Test2022_19_1_Sup2(t *testing.T) {
	data := `Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

	bestBlue := getBestBlue(data)

	if bestBlue != 12 {
		t.Errorf("Bad blueprints: %v (12)", bestBlue)
	}
}
