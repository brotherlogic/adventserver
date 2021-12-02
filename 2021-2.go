package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func procSub(str string) int {
	pos := 0
	depth := 0

	for _, st := range strings.Split(str, "\n") {
		elems := strings.Fields(st)
		dist, _ := strconv.Atoi(elems[1])
		switch elems[0] {
		case "forward":
			pos += dist
		case "down":
			depth += dist
		case "up":
			depth -= dist
		}

	}
	return pos * depth
}

func (s *Server) Solve2021day2part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(procSub(trimmed))}, nil
}
