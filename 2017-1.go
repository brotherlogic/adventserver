package main

import (
	"strconv"

	"golang.org/x/net/context"
)

func computeDigs(str string) int32 {
	sum := 0
	for i, c := range str {
		if i < len(str)-1 && string(c) == string(str[i+1]) {
			val, _ := strconv.Atoi(string(c))
			sum += val
		}
	}

	if str[0] == str[len(str)-1] {
		val, _ := strconv.Atoi(string(str[0]))
		sum += val
	}

	return int32(sum)
}

func (s *Server) solve2017day1part1(ctx context.Context) (int32, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-1.txt")
	if err != nil {
		return -1, err
	}

	return computeDigs(data), nil
}
