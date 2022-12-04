package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countFullyOverlapping(data string) int {
	count := 0

	for _, line := range strings.Split(data, "\n") {
		if doesOverlap(line) {
			count++
		}
	}

	return count
}

func countOverlapping(data string) int {
	count := 0

	for _, line := range strings.Split(data, "\n") {
		if doesOverlapEven(line) {
			count++
		}
	}

	return count
}

func getNums(piece string) []int {
	elems := strings.Split(piece, "-")
	num1, _ := strconv.ParseInt(elems[0], 10, 32)
	num2, _ := strconv.ParseInt(elems[1], 10, 32)

	return []int{int(num1), int(num2)}
}

func doesOverlap(line string) bool {
	if len(strings.TrimSpace(line)) == 0 {
		return false
	}
	elems := strings.Split(line, ",")

	nums1 := getNums(elems[0])
	nums2 := getNums(elems[1])

	if (nums1[0] <= nums2[0] && nums1[1] >= nums2[1]) ||
		(nums2[0] <= nums1[0] && nums2[1] >= nums1[1]) {
		return true
	}

	return false
}

func doesOverlapEven(line string) bool {
	if len(strings.TrimSpace(line)) == 0 {
		return false
	}
	elems := strings.Split(line, ",")

	nums1 := getNums(elems[0])
	nums2 := getNums(elems[1])

	if nums1[1] < nums2[0] || nums1[0] > nums2[1] {
		return false
	}

	return true
}

func (s *Server) Solve2022day4part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-4.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countFullyOverlapping(data))}, nil
}

func (s *Server) Solve2022day4part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-4.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countOverlapping(data))}, nil
}
