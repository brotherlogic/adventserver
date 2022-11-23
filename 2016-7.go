package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func tlsSupport(in string) bool {
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
