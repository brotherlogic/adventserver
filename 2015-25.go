package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getTriangleIndex(row, col int) int {
	top := (col * (col + 1)) / 2
	for i := 0; i < row-1; i++ {
		top = top + (col + i)
	}
	return top
}

func convertCode(base, row, col int) int {
	nth := getTriangleIndex(row, col)

	currv := base
	for i := 1; i < nth; i++ {
		currv *= 252533
		currv = currv % 33554393
	}

	return currv
}

func (s *Server) Solve2015day25part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{BigAnswer: int64(convertCode(20151125, 2981, 3075))}, nil
}
