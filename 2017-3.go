package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func computeSpiral(n int) int {
	power := 0
	for i := 1; ; i += 2 {
		if n <= i*i {
			power = i
			break
		}
	}

	top := power * power
	bottom := power*power - power + 1
	for i := 0; i < 4; i++ {
		if n >= bottom && n <= top {
			break
		}
		top -= (power - 1)
		bottom -= (power - 1)
	}

	dist := abs((top+bottom)/2 - n)

	return dist + (power-1)/2
}

func (s *Server) Solve2017day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(computeSpiral(265149))}, nil
}
