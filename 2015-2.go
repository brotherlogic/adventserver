package main

import (
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func computeAmountOfPaper(strin string) int32 {
	elems := strings.Split(strin, "\n")
	total := 0
	for _, str := range elems {
		strs := strings.Split(str, "x")
		length, err := strconv.Atoi(strs[0])
		if err == nil {
			width, err := strconv.Atoi(strs[1])
			if err == nil {
				height, err := strconv.Atoi(strs[2])
				if err == nil {
					surface1 := 2 * length * width
					surface2 := 2 * width * height
					surface3 := 2 * height * length
					slack := Min(surface1, Min(surface2, surface3)) / 2
					total += surface1 + surface2 + surface3 + slack
				}
			}
		}
	}

	return int32(total)
}

func computeAmountOfRibbon(strin string) int32 {
	elems := strings.Split(strin, "\n")
	total := 0
	for _, str := range elems {
		strs := strings.Split(str, "x")
		length, err := strconv.Atoi(strs[0])
		if err == nil {
			width, err := strconv.Atoi(strs[1])
			if err == nil {
				height, err := strconv.Atoi(strs[2])
				if err == nil {
					arr := []int{length, width, height}
					sort.Ints(arr)
					total += arr[0]*2 + arr[1]*2 + length*width*height
				}
			}
		}
	}

	return int32(total)
}

func (s *Server) Solve2015day2part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: computeAmountOfPaper(trimmed)}, nil
}

func (s *Server) Solve2015day2part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: computeAmountOfRibbon(trimmed)}, nil
}
