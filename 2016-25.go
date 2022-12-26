package main

import (
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findProgram(data string) int {
	for a := 0; a < math.MaxInt; a++ {
		res := runToggleProgram(data, a)

		if strings.HasPrefix("010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101", res.output) {
			return a
		}

	}
	return 0
}

func (s *Server) Solve2016day25part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-25.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(findProgram(data))}, nil
}
