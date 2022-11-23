package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func tlsSupport(in string) bool {
	inBracks := false
	for i := 0; i < len(in)-4; i++ {
		if in[i] == '[' {
			inBracks = true
		} else if in[i] == ']' {
			inBracks = false
		} else if in[i] == in[i+3] && in[i+1] == in[i+2] && in[i] != in[i+1] {
			return !inBracks
		}
	}
	return false
}

func (s *Server) Solve2016day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-7.txt")
	if err != nil {
		return nil, err
	}

	count := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			if tlsSupport(line) {
				count++
			}
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
