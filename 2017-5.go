package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runCode(nums []int64) int {
	pointer := 0
	steps := 0

	for pointer >= 0 && pointer < len(nums) {
		steps++
		jump := nums[pointer]
		nums[pointer]++
		pointer += int(jump)
	}

	return steps
}

func computeJumps(bits string) int {
	var nums []int64
	for _, numstr := range strings.Split(bits, "\n") {
		num, _ := strconv.Atoi(numstr)
		nums = append(nums, int64(num))
	}

	return runCode(nums)
}

func (s *Server) Solve2017day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-5.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeJumps(data))}, nil
}
