package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countInc(str string) int32 {
	count := int32(0)

	strs := strings.Split(str, "\n")
	for i, s := range strs {
		if i > 0 {
			n1, _ := strconv.Atoi(s)
			n2, _ := strconv.Atoi(strs[i-1])

			if n1 > n2 {
				count++
			}
		}
	}

	return count
}

func (s *Server) Solve2021day1part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-1.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(countInc(trimmed))}, nil
}
