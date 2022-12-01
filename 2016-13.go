package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runMaze(key, x, y int) int {
	return 0
}

func (s *Server) Solve2016day13part1(ctx context.Context) (*pb.SolveResponse, error) {
	res := runMaze(1364, 31, 39)

	return &pb.SolveResponse{Answer: int32(res)}, nil
}
