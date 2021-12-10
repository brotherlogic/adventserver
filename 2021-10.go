package main

import (
	"fmt"
	"sort"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getFirstInvalid(data string) (string, string) {
	var stack []string

	for _, char := range data {
		match := false
		switch string(char) {
		case "[", "(", "{", "<":
			stack = append(stack, string(char))
			continue

		case "]":
			match = stack[len(stack)-1] == "["
		case "}":
			match = stack[len(stack)-1] == "{"
		case ">":
			match = stack[len(stack)-1] == "<"
		case ")":
			match = stack[len(stack)-1] == "("
		}

		if match {
			stack = stack[0 : len(stack)-1]
		} else {
			return string(char), ""
		}
	}

	ret := ""
	for _, char := range stack {
		switch char {
		case "[":
			ret = "]" + ret
		case "{":
			ret = "}" + ret
		case "<":
			ret = ">" + ret
		case "(":
			ret = ")" + ret
		}
	}
	return "", ret
}

func getSum(data string) int {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		val, _ := getFirstInvalid(strings.TrimSpace(line))
		switch val {
		case ")":
			sum += 3
		case "]":
			sum += 57
		case "}":
			sum += 1197
		case ">":
			sum += 25137
		}
	}

	return sum
}
func (s *Server) getSum2(data string) int64 {
	var sums []int64
	for _, line := range strings.Split(data, "\n") {
		sum := int64(0)
		v1, val := getFirstInvalid(strings.TrimSpace(line))
		for _, c := range val {
			switch string(c) {
			case ")":
				sum = sum*5 + 1
			case "]":
				sum += sum*5 + 2
			case "}":
				sum += sum*5 + 3
			case ">":
				sum += sum*5 + 4
			}
		}

		if v1 == "" {
			sums = append(sums, sum)
		}
	}

	sort.SliceStable(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	s.Log(fmt.Sprintf("SUMS = %v", sums))

	return sums[len(sums)/2]
}

func (s *Server) Solve2021day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-10.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(getSum(trimmed))}, nil
}

func (s *Server) Solve2021day10part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-10.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{BigAnswer: s.getSum2(trimmed)}, nil
}
