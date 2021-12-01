package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countInc(str string) int32 {
	count := int32(0)

	strs := strings.Split(str, "\n")
	for i, s := range strs {
		if i > 0 {
			n1, _ := strconv.Atoi(s)
			n2, _ := strconv.Atoi(strs[i-1])

			if n1 > n2 {
				count++
			}
		}
	}

	return count
}

func countInc3(str string) int32 {
	var threes []int

	strs := strings.Split(str, "\n")
	for i, s := range strs {
		if i > 1 {
			n1, _ := strconv.Atoi(s)
			n2, _ := strconv.Atoi(strs[i-1])
			n3, _ := strconv.Atoi(strs[i-2])

			threes = append(threes, n1+n2+n3)
		}
	}

	str2 := ""
	for _, th := range threes {
		if len(str2) == 0 {
			str2 = fmt.Sprintf("%v", th)
		} else {
			str2 += fmt.Sprintf("\n%v", th)
		}
	}

	return countInc(str2)
}

func (s *Server) Solve2021day1part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-1.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(countInc(trimmed))}, nil
}

func (s *Server) Solve2021day1part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-1.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(countInc3(trimmed))}, nil
}
