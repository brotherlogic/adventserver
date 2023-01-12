package main

import (
	"log"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/adventserver/proto"
)

func computeSteps(data string) int {
	x, y := 0, 0

	for _, elem := range strings.Split(data, ",") {
		switch elem {
		case "n":
			y += 2
		case "s":
			y -= 2
		case "ne":
			x++
			y++
		case "se":
			x++
			y--
		case "nw":
			x--
			y++
		case "sw":
			x--
			y--
		default:
			log.Fatalf("Unknown direction: %v", elem)
		}
	}

	// Move to a diag
	if x == 0 && y == 0 {
		return 0
	}

	// We are below a diagnol
	if abs(y) > abs(x) {
		return (abs(y)-abs(x))/2 + abs(x)
	}

	if abs(y) == abs(x) {
		return abs(x)
	}

	return 0
}

func (s *Server) Solve2017day11part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-11.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeSteps(strings.TrimSpace(data)))}, nil
}
