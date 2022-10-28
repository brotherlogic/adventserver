package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildCubeArr(width int) [][][]bool {
	var ans [][][]bool

	for i := 0; i < width*2+1; i++ {
		var an [][]bool
		for j := 0; j < width*2+1; j++ {
			an = append(an, make([]bool, width*2+1))
		}

		ans = append(ans, an)
	}

	return ans
}

func doSplit(str string) (int, int) {
	elems := strings.Split(str, "=")
	bits := strings.Split(elems[1], "..")

	low, _ := strconv.Atoi(bits[0])
	high, _ := strconv.Atoi(bits[1])

	return low, high
}

func applyLine(line string, in [][][]bool) [][][]bool {
	elems := strings.Fields(line)

	adjust := false
	if elems[0] == "on" {
		adjust = true
	}

	nelems := strings.Split(elems[1], ",")
	lowx, highx := doSplit(nelems[0])
	if lowx+50 > 100 || highx+50 < 0 || lowx+50 < 0 {
		return in
	}
	for x := lowx; x <= highx; x++ {
		lowy, highy := doSplit(nelems[1])
		for y := lowy; y <= highy; y++ {
			lowz, highz := doSplit(nelems[2])
			for z := lowz; z <= highz; z++ {
				nx, ny, nz := x+50, y+50, z+50
				if nx >= 0 && nx <= 100 && ny >= 0 && ny <= 100 && nz >= 0 && nz <= 100 {
					in[nx][ny][nz] = adjust
				}
			}
		}
	}

	return in
}

func buildCubeAndCount(data string) int {
	arr := buildCubeArr(50)

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		arr = applyLine(strings.TrimSpace(line), arr)
	}

	count := 0
	for _, line := range arr {
		for _, nline := range line {
			for _, val := range nline {
				if val {
					count++
				}
			}
		}
	}
	return count
}

func (s *Server) Solve2021day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(buildCubeAndCount(data))}, nil
}
