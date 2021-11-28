package main

import (
	"log"
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func procCode(code string, sx, sy int) int32 {
	res := 0
	keypad := [3][3]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	elems := strings.Split(code, "\n")
	multiplier := int(math.Pow10(len(elems) - 1))
	for _, elem := range elems {
		for _, c := range elem {
			switch c {
			case 'U':
				sy++
			case 'D':
				sy--
			case 'L':
				sx--
			case 'R':
				sx++
			default:
				log.Fatalf("Cannot process %v", c)
			}

			if sx < 0 {
				sx = 0
			}
			if sx > 2 {
				sx = 2
			}
			if sy < 0 {
				sy = 0
			}
			if sy > 2 {
				sy = 2
			}
		}

		res += multiplier * keypad[sx][sy]
		multiplier /= 10
	}
	return int32(res)
}

func (s *Server) Solve2016day2part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: procCode(trimmed, 1, 1)}, nil
}
