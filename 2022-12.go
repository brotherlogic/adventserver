package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runMap(data string) int {
	return 0
}

func (s *Server) Solve2022day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-12.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runMap(data))}, nil
}
