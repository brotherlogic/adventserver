package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countVisibleTrees(data string) int {
	return 0
}

func (s *Server) Solve2022day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-8.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countVisibleTrees(data))}, nil
}
