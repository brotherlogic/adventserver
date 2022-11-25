package main

import "testing"

func Test2016_10_1(t *testing.T) {
	data := `value 5 goes to bot 2
	bot 2 gives low to bot 1 and high to bot 0
	value 3 goes to bot 1
	bot 1 gives low to output 1 and high to bot 0
	bot 0 gives low to output 2 and high to output 0
	value 2 goes to bot 2`

	found := false
	res := runBotProgram(data)
	for _, r := range res {
		if r.comp(2, 5) {
			found = true
		}
	}

	if !found {
		t.Errorf("Program did not run correctly: %+v", res)
	}
}
