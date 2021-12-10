package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getFirstInvalid(data string) string {
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
			return string(char)
		}
	}

	return ""
}

func getSum(data string) int {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		switch getFirstInvalid(strings.TrimSpace(line)) {
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

func (s *Server) Solve2021day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-10.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(getSum(trimmed))}, nil
}
