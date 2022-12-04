package main

func nextLine(line string) string {
	return ""
}

func countSafes(line string) int {
	return 0
}

func runTiles(line string, c int) int {
	safes := countSafes(line)

	for i := 0; i < c; i++ {
		line = nextLine(line)
		safes += countSafes(line)
	}

	return safes
}
