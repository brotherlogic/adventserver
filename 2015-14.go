package main

import (
	"context"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
)

func getDistance(data string, time int) int32 {
	return int32(0)
}

func (s *Server) Solve2015day14part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-14.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: getDistance(trimmed, 2503)}, nil
}
