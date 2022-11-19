package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findMaxHouse(val int) int {
	return 0
}

func (s *Server) Solve2015day20part1(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{Answer: int32(findMaxHouse(36000000))}, nil
}
