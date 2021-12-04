package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runCode(nums []int) int {
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

func runCode2(nums []int) int {
	pointer := 0
	steps := 0

	for pointer >= 0 && pointer < len(nums) {
		steps++
		jump := nums[pointer]

		if nums[pointer] > 2 {
			nums[pointer]--
		} else {
			nums[pointer]++
		}
		pointer += int(jump)
	}

	return steps
}

func computeJumps(bits string) int {
	var nums []int
	for _, numstr := range strings.Split(bits, "\n") {
		num, _ := strconv.Atoi(numstr)
		nums = append(nums, num)
	}

	return runCode(nums)
}

func computeJumps2(bits string) int {
	var nums []int
	for _, numstr := range strings.Split(bits, "\n") {
		num, _ := strconv.Atoi(numstr)
		nums = append(nums, num)
	}

	return runCode2(nums)
}

func (s *Server) Solve2017day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-5.txt")
	if err != nil {
		return nil, err
	}
	data = strings.TrimSpace(data)
	return &pb.SolveResponse{Answer: int32(computeJumps(data))}, nil
}

func (s *Server) Solve2017day5part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-5.txt")
	if err != nil {
		return nil, err
	}
	data = strings.TrimSpace(data)
	return &pb.SolveResponse{Answer: int32(computeJumps2(data))}, nil
}
