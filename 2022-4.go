package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countFullyOverlapping(data string) int {
	return 0
}

func doesOverlap(line string) bool {
	return false
}

func (s *Server) Solve2022da43part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-4.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countFullyOverlapping(data))}, nil
}
