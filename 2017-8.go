package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type registers struct {
	register map[string]int
}

func runJumpProgram(data string) *registers {
	registers := &registers{register: make(map[string]int)}
	return registers
}

func (s *Server) Solve2017day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-7.txt")
	if err != nil {
		return nil, err
	}
	res := runJumpProgram(data)

	highest := 0
	for _, value := range res.register {
		if value > highest {
			highest = value
		}
	}

	return &pb.SolveResponse{Answer: int32(highest)}, nil
}
