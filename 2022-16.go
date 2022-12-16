package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func releaseGas(data string, minutes int) int {
	return 0
}

func (s *Server) Solve2022day16part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-16.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(releaseGas(data, 30))}, nil
}
