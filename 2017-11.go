package main

import (
	"strings"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/adventserver/proto"
)

func computeSteps(data string) int {
	return 0
}

func (s *Server) Solve2017day11part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-11.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeSteps(strings.TrimSpace(data)))}, nil
}
