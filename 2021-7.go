package main

import (
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func computeMin(nums []int) int {
	mx := nums[0]
	mi := nums[0]
	for _, num := range nums {
		if num > mx {
			mx = num
		}

		if num < mi {
			mi = num
		}
	}

	best := math.MaxInt32
	for i := mi; i <= mx; i++ {
		sumv := 0
		for _, num := range nums {
			sumv += abs(num - i)
		}

		if sumv < best {
			best = sumv
		}
	}

	return best
}

func triangle(num int) int {
	if num == 0 {
		return 0
	}
	return num * (num + 1) / 2
}

func computeMinComplex(nums []int) int {
	mx := nums[0]
	mi := nums[0]
	for _, num := range nums {
		if num > mx {
			mx = num
		}

		if num < mi {
			mi = num
		}
	}

	best := math.MaxInt32
	for i := mi; i <= mx; i++ {
		sumv := 0
		for _, num := range nums {
			sumv += triangle(abs(num - i))
		}

		if sumv < best {
			best = (sumv)
		}
	}

	return best
}

func getCost(str string) int {
	var nums []int

	for _, elem := range strings.Split(str, ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(elem))
		nums = append(nums, n)
	}

	return computeMin(nums)
}

func getCostComplex(str string) int {
	var nums []int

	for _, elem := range strings.Split(str, ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(elem))
		nums = append(nums, n)
	}

	return computeMinComplex(nums)
}

func (s *Server) Solve2021day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-7a.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(getCost(trimmed))}, nil
}

func (s *Server) Solve2021day7part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-7a.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(getCostComplex(trimmed))}, nil
}
