package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runLightProgram(x, y int, program string) [][]bool {
	b := make([][]bool, y)
	for i := range b {
		b[i] = make([]bool, x)
	}

	return b
}

func countBoolArr(b [][]bool) int {
	count := 0
	for x := range b {
		for y := range b[x] {
			if b[x][y] {
				count++
			}
		}
	}
	return count
}

func (s *Server) Solve2016day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-8.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countBoolArr(runLightProgram(50, 6, data)))}, nil
}
