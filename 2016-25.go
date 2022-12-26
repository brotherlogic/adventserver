package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findProgram(data string) int {
	a := 1
	res := runToggleProgram(data, a)

	if res.output == "01010101" {
		return a
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
