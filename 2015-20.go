package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findMaxHouse(sval int) int {
	houses := make(map[int]int)

	start := 1
	for {
		for i := 1; i <= 10; i++ {
			houses[start*i] += start * 10
		}

		for key, val := range houses {
			if val >= sval {
				return key
			}
		}
		start++
	}
}

func (s *Server) Solve2015day20part1(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{Answer: int32(findMaxHouse(36000000))}, nil
}
