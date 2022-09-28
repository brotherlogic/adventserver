package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"

	"golang.org/x/net/context"
)

func computeStringLength(str string) (int, int) {
	unicodes := strings.Count(str, `\x`)
	backslash := strings.Count(str, `\`)

	return len(str), len(str) - (unicodes * 3) - (backslash - unicodes) - 2
}

func (s *Server) Solve2015day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-8.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	chs := 0
	cds := 0
	for _, str := range strings.Split(trimmed, "\n") {
		c1, c2 := computeStringLength(str)
		chs += c1
		cds += c2
	}

	return &pb.SolveResponse{Answer: int32(chs - cds)}, nil
}
