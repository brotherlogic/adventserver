package main

import (
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countCalories(data string) int {

	best := 0

	sofar := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			if sofar > best {
				best = sofar
			}
			sofar = 0
		} else {
			val, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
			sofar += int(val)
		}
	}

	return best
}

func topThreeCalories(data string) int {

	sofar := 0

	var amounts []int
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			amounts = append(amounts, sofar)
			sofar = 0
		} else {
			val, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
			sofar += int(val)
		}
	}
	amounts = append(amounts, sofar)

	sort.Sort(sort.Reverse(sort.IntSlice(amounts)))

	return amounts[0] + amounts[1] + amounts[2]
}

func (s *Server) Solve2022day1part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-1.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countCalories(data))}, nil
}

func (s *Server) Solve2022day1part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-1.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(topThreeCalories(data))}, nil
}
