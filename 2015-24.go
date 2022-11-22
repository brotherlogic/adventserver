package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func computeGrouping(weights []int) int {
	return 0
}

func (s *Server) Solve2015day24part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-24.txt")
	if err != nil {
		return nil, err
	}

	var weights []int
	for _, line := range strings.Split(data, "\n") {
		val, _ := strconv.ParseInt(line, 10, 32)
		weights = append(weights, int(val))
	}

	return &pb.SolveResponse{Answer: int32(computeGrouping(weights))}, nil
}
