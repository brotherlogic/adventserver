package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func nextLine(line string) string {
	newLine := ""
	nline := "." + line + "."
	for i := 1; i < len(nline)-1; i++ {
		if (nline[i-1] == '^' && nline[i] == '^' && nline[i+1] == '.') ||
			(nline[i-1] == '.' && nline[i] == '^' && nline[i+1] == '^') ||
			(nline[i-1] == '^' && nline[i] == '.' && nline[i+1] == '.') ||
			(nline[i-1] == '.' && nline[i] == '.' && nline[i+1] == '^') {
			newLine += "^"
		} else {
			newLine += "."
		}
	}
	return newLine
}

func countSafes(line string) int {
	c := 0
	for _, rune := range line {
		if rune == '.' {
			c++
		}
	}

	return c
}

func runTiles(line string, c int) int {
	safes := countSafes(line)

	for i := 0; i < c-1; i++ {
		line = nextLine(line)
		safes += countSafes(line)
	}

	return safes
}

func (s *Server) Solve2016day18part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runTiles("^^.^..^.....^..^..^^...^^.^....^^^.^.^^....^.^^^...^^^^.^^^^.^..^^^^.^^.^.^.^.^.^^...^^..^^^..^.^^^^", 40))}, nil
}
