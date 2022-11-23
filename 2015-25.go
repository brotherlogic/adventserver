package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func convertCode(base, row, col int) int {
	return 0
}

func (s *Server) Solve2015day25part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{BigAnswer: int64(convertCode(20151125, 2981, 3075))}, nil
}
