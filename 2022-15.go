package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countKnown(data string, y int) int {
	return 0
}

func (s *Server) Solve2022day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-15.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countKnown(data, 2000000))}, nil
}
