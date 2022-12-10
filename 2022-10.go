package main

type elfProgram struct {
	a      int
	cycle  int
	values []int
}

func (e elfProgram) getSignal() int {
	return 0
}

func runElfProgram(data string) elfProgram {
	return elfProgram{}
}
