package main

type bot struct {
	low, high int
	num       int
	comps     [][]int
}

func (b bot) comp(val1, val2 int) bool {
	for _, c := range b.comps {
		if c[0] == val1 && c[1] == val2 || c[0] == val2 && c[1] == val1 {
			return true
		}
	}
	return false
}

func runBotProgram(data string) []bot {
	var bots []bot

	return bots
}
