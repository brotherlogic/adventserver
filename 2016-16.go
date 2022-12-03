package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func dragonExpand(in string) string {
	return ""
}

func dragonRun(in string, size int) string {
	return ""
}

func (s *Server) Solve2016day16part1(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{StringAnswer: (dragonRun("11100010111110100", 272))}, nil
}
