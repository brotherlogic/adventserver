package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countZeros(data string) int {

	connections := make(map[int][]int)

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		bits := strings.Split(line, "<->")
		num1 := getInt32(bits[0])
		if _, ok := connections[num1]; !ok {
			connections[num1] = make([]int, 0)
		}

		for _, bit := range strings.Split(bits[1], ",") {
			num2 := getInt32(strings.TrimSpace(bit))
			if _, ok := connections[num2]; !ok {
				connections[num2] = make([]int, 0)
			}

			connections[num1] = append(connections[num1], num2)
			connections[num2] = append(connections[num2], num1)
		}
	}

	seen := make(map[int]bool)
	queue := []int{0}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		seen[head] = true

		for _, next := range connections[head] {
			if !seen[next] {
				queue = append(queue, next)
			}
		}
	}

	return len(seen)
}

func (s *Server) Solve2017day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-12.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countZeros(data))}, nil
}
