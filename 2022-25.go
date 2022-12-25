package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func snafu(in string) int {
	return 0
}

func rsnafu(in int) string {
	return ""
}

func computeSnafuSum(data string) string {
	sumv := 0
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		sumv += snafu(strings.TrimSpace(line))
	}

	return rsnafu(sumv)
}

func (s *Server) Solve2022day25part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-25.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: computeSnafuSum(data)}, nil
}
