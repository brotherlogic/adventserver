package main

import (
	"strings"

	"golang.org/x/net/context"
)

func ComputeFloor(str string) int {
	left := strings.Count(str, "(")
	right := strings.Count(str, ")")
	return left - right
}

func ComputeF1(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			count++
		} else if str[i] == ')' {
			count--
		}

		if count == -1 {
			return i + 1
		}
	}

	return -1
}

func (s *Server) solve2015day1part1(ctx context.Context) (int32, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-1.txt")
	if err != nil {
		return -1, err
	}

	return int32(ComputeFloor(data)), nil
}

func (s *Server) solve2015day1part2(ctx context.Context) (int32, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-1.txt")
	if err != nil {
		return -1, err
	}

	return int32(ComputeF1(data)), nil
}
