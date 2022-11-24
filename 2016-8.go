package main

func runLightProgram(x, y int, program string) [][]bool {
	b := make([][]bool, y)
	for i := range b {
		b[i] = make([]bool, x)
	}

	return b
}
