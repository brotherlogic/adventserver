package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func rearrangeCrates(data string) string {
	return ""
}

func (s *Server) Solve2022day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-5.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: rearrangeCrates(strings.TrimSpace(data))}, nil
}
