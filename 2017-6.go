package main

import (
	"fmt"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func doNext(blocks []int) []int {
	maxV := 0
	maxI := 0

	for i := range blocks {
		if blocks[i] > maxV {
			maxV = blocks[i]
			maxI = i
		}
	}

	val := blocks[maxI]
	blocks[maxI] = 0
	for i := 1; i <= val; i++ {
		blocks[(maxI+i)%len(blocks)]++
	}

	return blocks
}

func computeRepeat(blocks []int) (int, int) {
	seen := make(map[string]int)
	count := 0
	seen[fmt.Sprintf("%v", blocks)] = 1

	for {
		blocks = doNext(blocks)

		if seen[fmt.Sprintf("%v", blocks)] > 0 {
			return count + 1, count - seen[fmt.Sprintf("%v", blocks)] + 1
		} else {
			count++
			seen[fmt.Sprintf("%v", blocks)] = count
		}
	}
}

func (s *Server) Solve2017day6part1(ctx context.Context) (*pb.SolveResponse, error) {
	a1, _ := computeRepeat([]int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5})
	return &pb.SolveResponse{Answer: int32(a1)}, nil
}

func (s *Server) Solve2017day6part2(ctx context.Context) (*pb.SolveResponse, error) {
	_, a1 := computeRepeat([]int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5})
	return &pb.SolveResponse{Answer: int32(a1)}, nil
}
