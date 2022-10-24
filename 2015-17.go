package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildArrInt(data string) []int {
	var res []int
	for _, line := range strings.Split(data, "\n") {
		nv, _ := strconv.ParseInt(line, 10, 64)
		res = append(res, int(nv))
	}
	return res
}

func computeContainers(data string, total int) int {
	arr := buildArrInt(data)

	result := doContainers(arr, total, 0)
	return result
}

func doContainers(arr []int, total, sofar int) int {
	return 0
}

func (s *Server) Solve2015day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-17.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeContainers(trimmed, 150))}, nil
}
