package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getShortestPath(key string) string {
	return ""
}

func (s *Server) Solve2016day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{StringAnswer: (getShortestPath("mmsxrhfx"))}, nil
}
