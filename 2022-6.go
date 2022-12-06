package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findMarker(str string) int {
	return 0
}

func (s *Server) Solve2022day6part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-6.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(findMarker(data))}, nil
}
