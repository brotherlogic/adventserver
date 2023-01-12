package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countZeros(data string) int {
	return 0
}

func (s *Server) Solve2017day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-12.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countZeros(data))}, nil
}
