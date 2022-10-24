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
		nv, _ := strconv.ParseInt(line, 10, 32)
		res = append(res, int(nv))
	}
	return res
}

func computeContainers(data string, total int) int {
	arr := buildArrInt(data)

	result := doContainers(arr, total, 0, 0)
	sumv := 0
	for _, v := range result {
		sumv += v
	}

	return sumv
}

func computeMinContainers(data string, total int) int {
	arr := buildArrInt(data)

	result := doContainers(arr, total, 0, 0)
	sumv := 0
	minv := 2000
	for k, v := range result {
		if k < minv {
			minv = k
			sumv = v
		}
	}

	return sumv
}

func doContainers(arr []int, total, sofar, c int) map[int]int {
	if sofar == total {
		return map[int]int{c: 1}
	}

	if sofar > total {
		return make(map[int]int)
	}

	mmap := make(map[int]int)
	for i, p := range arr {
		var narr []int
		for _, val := range arr[i+1:] {
			narr = append(narr, val)
		}
		mapper := doContainers(narr, total, sofar+p, c+1)
		for k, v := range mapper {
			mmap[k] += v
		}
	}

	return mmap
}

func (s *Server) Solve2015day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-17.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeContainers(trimmed, 150))}, nil
}

func (s *Server) Solve2015day17part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-17.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeMinContainers(trimmed, 150))}, nil
}
