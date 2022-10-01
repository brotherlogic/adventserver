package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func lookAndSay(s string) string {
	return s
}

func (s *Server) Solve2015day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	start := "3113322113"
	for i := 0; i < 40; i++ {
		start = lookAndSay(start)
	}

	return &pb.SolveResponse{BigAnswer: int64(len(start))}, nil
}
