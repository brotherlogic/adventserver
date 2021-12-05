package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func build(data string) (int, int) {
	elems := strings.Split(data, ",")

	a, _ := strconv.Atoi(strings.TrimSpace(elems[0]))
	b, _ := strconv.Atoi(strings.TrimSpace(elems[1]))

	return a, b
}

func buildCoord(data string) (int, int, int, int) {
	bits := strings.Split(data, "->")

	a, b := build(bits[0])
	c, d := build(bits[1])

	return a, b, c, d
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeCrosses(data string) int {
	var smap [][]int

	for i := 0; i < 10; i++ {
		smap = append(smap, make([]int, 10))
	}

	for _, coord := range strings.Split(data, "\n") {
		x1, y1, x2, y2 := buildCoord(coord)

		if x1 == x2 {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				smap[x1][y]++
			}
		}

		if y1 == y2 {
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				smap[x][y1]++
			}
		}
	}

	count := 0
	for _, c := range smap {
		for _, cc := range c {
			if cc > 1 {
				count++
			}
		}
	}

	return count
}

func (s *Server) Solve2021day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-5.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeCrosses(data))}, nil
}
