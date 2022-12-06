package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findMarker(str string, cc int) int {
	for i := 0; i < len(str)-3; i++ {
		counts := make(map[byte]bool)
		found := false
		for j := i; j < i+cc; j++ {
			if _, ok := counts[str[j]]; ok {
				found = true
				break
			} else {
				counts[str[j]] = true
			}
		}

		if !found {
			return i + cc
		}
	}

	return 0
}

func (s *Server) Solve2022day6part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-6.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(findMarker(data, 4))}, nil
}

func (s *Server) Solve2022day6part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-6.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(findMarker(data, 14))}, nil
}
