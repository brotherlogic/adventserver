package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type mstate struct {
	a, b, c, d int
}

func runMonorailProgram(data string) mstate {
	return mstate{}
}

func (s *Server) Solve2016day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-12.txt")
	if err != nil {
		return nil, err
	}

	state := runMonorailProgram(data)

	return &pb.SolveResponse{Answer: int32(state.a)}, nil
}
