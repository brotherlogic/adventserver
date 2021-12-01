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
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < len(values) && i >= 0 && j < len(values[i]) && j >= 0 {
				sumv += values[i][j]
			}
		}
	}
	return sumv
}

func buildSpiral(n int) int {
	var values [][]int
	width := 50
	for i := 0; i < width*2; i++ {
		values = append(values, make([]int, width*2))
	}

	spx := width
	spy := width

	values[spx][spy] = 1
	for w := 1; w < width; w++ {
		nspx := width + w
		nspy := width - w + 1

		dir := 0
		for count := 0; count < w*8; count++ {
			value := computeSpiralPoint(values, nspx, nspy)
			if value > n {
				return value
			}
			values[nspx][nspy] = value
			switch dir {
			case 0:
				nspy++
				if nspy-width == w {
					dir = 1
				}
			case 1:
				nspx--
				if width-nspx == w {
					dir = 2
				}
			case 2:
				nspy--
				if width-nspy == w {
					dir = 3
				}
			case 3:
				nspx++
			}
		}
	}

	return 0
}

func (s *Server) Solve2017day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(computeSpiral(265149))}, nil
}

func (s *Server) Solve2017day3part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(buildSpiral(265149))}, nil
}
