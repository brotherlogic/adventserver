package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runTetris(data string) []int {
	return []int{0}
}

func getHeight(h []int) int {
	highest := 0
	for _, value := range h {
		if value > highest {
			highest = value
		}
	}

	return highest
}

func (s *Server) Solve2022day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-17.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getHeight(runTetris(data)))}, nil
}
