package main

import "testing"

func Test2021Day14(t *testing.T) {
	data := `NNCB

	CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C`

	newone := runData(data, 10)
	mc, lc := getCommons(newone)
	if mc-lc != 1588 {
		t.Errorf("Bad day: %v", mc-lc)
	}
}

func Test2021Day14Focus(t *testing.T) {
	rules := `CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C`

	res := runRules("NNCB", buildRules(rules))
	if res != "NCNBCHB" {
		t.Errorf("Bad res: %v", res)
	}

	res = runRules("NCNBCHB", buildRules(rules))
	if res != "NBCCNBBBCBHCB" {
		t.Errorf("Bad res: %v", res)
	}

	res = runRules("NBBBCNCCNBBNBNBBCHBHHBCHB", buildRules(rules))
	if res != "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB" {
		t.Errorf("Bad res: %v", res)
	}
}
