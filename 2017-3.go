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

func computeSpiralPoint(values [][]int, x, y int) int {
	sumv := 0
	for i := -1; i <= 1; i++ {
		for j := 1; j <= 1; j++ {
			if i < len(values) && i >= 0 && j < len(values[i]) && j >= 0 {
				sumv += values[i][j]
			}
		}
	}
	return sumv
}

func buildSpiral(n int) int {
	var values [][]int
	for i := 0; i < 100; i++ {
		values = append(values, make([]int, 100))
	}

	spx := 50
	spy := 50

	values[spx][spy] = 1

	return 0
}

func (s *Server) Solve2017day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(computeSpiral(265149))}, nil
}
