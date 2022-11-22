package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type computer struct {
	a, b int
}

func runProgram(program string) computer {
	return computer{}
}

func (s *Server) Solve2015day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-23.txt")
	if err != nil {
		return nil, err
	}

	result := runProgram(data)
	return &pb.SolveResponse{Answer: int32(result.a)}, nil
}
