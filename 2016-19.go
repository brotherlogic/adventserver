package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runPresents(num int) int {
	return 0
}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
