package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runArray(list, steps []int) int32 {
	return 0
}

func (s *Server) Solve2017day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	var list []int
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	return &pb.SolveResponse{Answer: runArray(list, []int{212, 254, 178, 237, 2, 0, 1, 54, 167, 92, 117, 125, 255, 61, 159, 164})}, nil
}
