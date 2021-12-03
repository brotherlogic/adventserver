package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func isValidPassword(in string) bool {
	counts := make(map[string]int)
	for _, word := range strings.Fields(in) {
		counts[word]++
	}

	for _, val := range counts {
		if val > 1 {
			return false
		}
	}

	return true
}

func (s *Server) Solve2017day4part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-d4.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	count := 0
	for _, str := range strings.Split(trimmed, "\n") {
		if isValidPassword(str) {
			count++
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
