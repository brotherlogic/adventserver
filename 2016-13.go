package main

import (
	"strconv"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	mapPoints = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day13_points",
		Help: "The number of server requests",
	})
)

type mazePoint struct {
	x, y int
	best int
}

func isWall(key, x, y int) bool {
	num := x*x + 3*x + 2*x*y + y + y*y
	num += key
	binary := strconv.FormatInt(int64(num), 2)
	count := 0
	for _, char := range binary {
		if char == '1' {
			count++
		}
	}

	return count%2 == 1
}

func isLegitMove(key int, seen map[int]map[int]int, x, y int) bool {
	if isWall(key, x, y) {
		return false
	}

	if val, ok := seen[x]; ok {
		if _, aok := val[y]; aok {
			return false
		}
	}

	return true
}

func runMaze(key, x, y int) int {
	var queue []*mazePoint
	queue = append(queue, &mazePoint{x: 1, y: 1, best: 0})

	seen := make(map[int]map[int]int)
	seen[1] = make(map[int]int)
	seen[1][1] = 0

	for {
		mapPoints.Inc()
		head := queue[0]
		queue = queue[1:]

		// Look for a win
		if head.x == x && head.y == y {
			return head.best
		}

		// Generate all other locations
		if isLegitMove(key, seen, head.x+1, head.y) {
			queue = append(queue, &mazePoint{head.x + 1, head.y, head.best + 1})
			if _, ok := seen[head.x+1]; !ok {
				seen[head.x+1] = make(map[int]int)
			}
			seen[head.x+1][head.y] = head.best + 1
		}

		if isLegitMove(key, seen, head.x-1, head.y) {
			queue = append(queue, &mazePoint{head.x - 1, head.y, head.best + 1})
			if _, ok := seen[head.x-1]; !ok {
				seen[head.x-1] = make(map[int]int)
			}
			seen[head.x-1][head.y] = head.best + 1

		}

		if isLegitMove(key, seen, head.x, head.y+1) {
			queue = append(queue, &mazePoint{head.x, head.y + 1, head.best + 1})
			seen[head.x][head.y+1] = head.best + 1
		}

		if isLegitMove(key, seen, head.x, head.y-1) {
			queue = append(queue, &mazePoint{head.x, head.y - 1, head.best + 1})
			seen[head.x][head.y-1] = head.best + 1
		}
	}

}

func runMazeToLimit(limit, key, x, y int) int {
	var queue []*mazePoint
	queue = append(queue, &mazePoint{x: 1, y: 1, best: 0})

	seen := make(map[int]map[int]int)
	seen[1] = make(map[int]int)
	seen[1][1] = 0

	for {
		mapPoints.Inc()
		head := queue[0]
		queue = queue[1:]

		// Look for a win
		if head.best > limit {
			count := 0
			for _, val := range seen {
				for _, vval := range val {
					if vval == limit {
						count++
					}
				}
			}
		}

		// Generate all other locations
		if isLegitMove(key, seen, head.x+1, head.y) {
			queue = append(queue, &mazePoint{head.x + 1, head.y, head.best + 1})
			if _, ok := seen[head.x+1]; !ok {
				seen[head.x+1] = make(map[int]int)
			}
			seen[head.x+1][head.y] = head.best + 1
		}

		if isLegitMove(key, seen, head.x-1, head.y) {
			queue = append(queue, &mazePoint{head.x - 1, head.y, head.best + 1})
			if _, ok := seen[head.x-1]; !ok {
				seen[head.x-1] = make(map[int]int)
			}
			seen[head.x-1][head.y] = head.best + 1

		}

		if isLegitMove(key, seen, head.x, head.y+1) {
			queue = append(queue, &mazePoint{head.x, head.y + 1, head.best + 1})
			seen[head.x][head.y+1] = head.best + 1
		}

		if isLegitMove(key, seen, head.x, head.y-1) {
			queue = append(queue, &mazePoint{head.x, head.y - 1, head.best + 1})
			seen[head.x][head.y-1] = head.best + 1
		}
	}

}

func (s *Server) Solve2016day13part1(ctx context.Context) (*pb.SolveResponse, error) {
	res := runMaze(1364, 31, 39)

	return &pb.SolveResponse{Answer: int32(res)}, nil
}

func (s *Server) Solve2016day13part2(ctx context.Context) (*pb.SolveResponse, error) {
	res := runMazeToLimit(50, 1364, 31, 39)

	return &pb.SolveResponse{Answer: int32(res)}, nil
}
