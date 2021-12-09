package main

import "testing"

func Test2017Part7(t *testing.T) {
	data := `pbga (66)
	xhth (57)
	ebii (61)
	havc (66)
	ktlj (57)
	fwft (72) -> ktlj, cntj, xhth
	qoyq (66)
	padx (45) -> pbga, havc, qoyq
	tknk (41) -> ugml, padx, fwft
	jptl (61)
	ugml (68) -> gyxo, ebii, jptl
	gyxo (61)
	cntj (57)`

	bot := getBottom(data)
	if bot != "tknk" {
		t.Errorf("Bottom is wrong %v vs tknk", bot)
	}
}

func Test2017Check(t *testing.T) {
	data := `czlmv (78)
	pupaehu (99) -> eolilw, czlmv, zlrvs, vrppl`

	bot := getBottom(data)
	if bot != "pupaehu" {
		t.Errorf("Bah: %v", bot)
	}
}
