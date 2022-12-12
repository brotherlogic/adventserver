package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type toggler struct {
	a       int
	program []string
	pointer int
}

func runToggleProgram(data string) *toggler {
	toggler := &toggler{}

	return toggler
}

func (s *Server) Solve2016day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-23.txt")
	if err != nil {
		return nil, err
	}

	res := runToggleProgram(data)
	return &pb.SolveResponse{Answer: int32(res.a)}, nil
}
