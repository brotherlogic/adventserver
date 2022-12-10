package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type elfProgram struct {
	a      int
	cycle  int
	values []int
}

func (e elfProgram) getSignal() int {
	return 0
}

func runElfProgram(data string) elfProgram {
	return elfProgram{}
}


func (s *Server) Solve2022day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-10.txt")
	if err != nil {
		return nil, err
	}

	result := runElfProgram(data)

	return &pb.SolveResponse{Answer: int32(result.getSignal())}, nil
}