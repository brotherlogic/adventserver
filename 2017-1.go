package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func computeDigs(str string) int32 {
	sum := 0
	for i, c := range str {
		if i < len(str)-1 && string(c) == string(str[i+1]) {
			val, _ := strconv.Atoi(string(c))
			sum += val
		}
	}

	if str[0] == str[len(str)-1] {
		val, _ := strconv.Atoi(string(str[0]))
		sum += val
	}

	return int32(sum)
}

func computeJump(str string) int32 {
	sum := 0
	for i, c := range str {
		if string(c) == string(str[(i+len(str)/2)%len(str)]) {
			val, _ := strconv.Atoi(string(c))
			sum += val
		}
	}

	return int32(sum)
}

func (s *Server) Solve2017day1part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-1.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: computeDigs(trimmed)}, nil
}

func (s *Server) Solve2017day1part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-1.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: computeJump(trimmed)}, nil
}
