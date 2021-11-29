package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func ComputeNumberOfHouses(str string) int {
	x := 0
	y := 0

	var m map[string]int
	m = make(map[string]int)

	m[strconv.Itoa(x)+"|"+strconv.Itoa(y)] = 1

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		default:
			panic("Unknown string")
		}

		m[strconv.Itoa(x)+"|"+strconv.Itoa(y)] = 1
	}

	return len(m)
}

func ComputeNumberOfRoboHouses(str string) int {
	hx := 0
	hy := 0
	rx := 0
	ry := 0

	var m map[string]int
	m = make(map[string]int)

	m[strconv.Itoa(hx)+"|"+strconv.Itoa(hy)] = 1

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '<':
			if i%2 == 0 {
				hx--
			} else {
				rx--
			}
		case '>':
			if i%2 == 0 {
				hx++
			} else {
				rx++
			}
		case '^':
			if i%2 == 0 {
				hy++
			} else {
				ry++
			}

		case 'v':
			if i%2 == 0 {
				hy--
			} else {
				ry--
			}

		default:
			panic("Unknown string")
		}

		m[strconv.Itoa(hx)+"|"+strconv.Itoa(hy)] = 1
		m[strconv.Itoa(rx)+"|"+strconv.Itoa(ry)] = 1
	}

	return len(m)
}

func (s *Server) Solve2015day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-3.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(ComputeNumberOfHouses(trimmed))}, nil
}

func (s *Server) Solve2015day3part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-3.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(ComputeNumberOfRoboHouses(trimmed))}, nil
}
