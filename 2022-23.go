package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runElves(data string, rounds int) int {
	return 0
}

func (s *Server) Solve2022day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-23.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runElves(data, 10))}, nil
}
