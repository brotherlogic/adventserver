package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type player struct {
	hitp, armor, damage int
}

func fight(p1, p2 player) bool {
	return false
}

func runFight(hitp, damage, armor int) int {
	return 0
}

func (s *Server) Solve2015day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runFight(103, 9, 2))}, nil
}
