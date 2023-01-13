package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	zlen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2017_12_qlen",
		Help: "The number of server requests",
	})
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
		zlen.Set(float64(len(queue)))
		head := queue[0]
		queue = queue[1:]

		for _, next := range connections[head] {
			if !seen[next] {
				queue = append(queue, next)
				seen[next] = true
			}
		}
	}

	return len(seen)
}

func countGroups(data string) int {

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

	scount := 1
	seen := make(map[int]int)

	for num := range connections {
		if seen[num] == 0 {
			queue := []int{num}

			for len(queue) > 0 {
				zlen.Set(float64(len(queue)))
				head := queue[0]
				queue = queue[1:]

				for _, next := range connections[head] {
					if seen[next] == 0 {
						queue = append(queue, next)
						seen[next] = scount
					}
				}
			}
			scount++
		}

	}

	return scount - 1
}

func (s *Server) Solve2017day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-12.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countZeros(data))}, nil
}

func (s *Server) Solve2017day12part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-12.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countGroups(data))}, nil
}
