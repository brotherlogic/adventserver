package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runCycle(mapper map[int]int) map[int]int {
	nmapper := make(map[int]int)
	for key, count := range mapper {
		if key == 0 {
			nmapper[6] += count
			nmapper[8] += count
		} else {
			nmapper[key-1] += count
		}
	}

	return nmapper
}

func computeCycle(list string, days int) int {
	mapper := make(map[int]int)

	for _, elem := range strings.Split(list, ",") {
		val, _ := strconv.Atoi(elem)
		mapper[val]++
	}

	for i := 0; i < days; i++ {
		mapper = runCycle(mapper)
	}

	count := 0
	for _, c := range mapper {
		count += c
	}
	return count
}

func (s *Server) Solve2021day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-7.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeCycle(strings.TrimSpace(data), 80))}, nil
}
