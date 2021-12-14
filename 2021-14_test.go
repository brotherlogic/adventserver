package main

import (
	"log"
	"testing"
)

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
	if mc.Sub(mc, lc).String() != "1588" {
		t.Fatalf("Bad day: %v", mc.Sub(mc, lc).String())
	}

	newone = runData(data, 40)
	mc, lc = getCommons(newone)
	if mc.Sub(mc, lc).String() != "2188189693529" {
		t.Errorf("Bad 2nd day: %v", mc.Sub(mc, lc).String())
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

	res := runRules(convertToMap("NNCB"), buildRules(rules))
	log.Printf("%v", res)
}
