package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	evals = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day23_evals",
		Help: "The number of server requests",
	})
)

func computeGrouping(weights []int, count int) int {

	sumv := 0
	for _, weight := range weights {
		sumv += weight
	}
	goal := sumv / count

	res := altGrouping(weights, goal)

	return res
}

func prod(v []int) int {
	prodv := 1
	for _, val := range v {
		prodv *= val
	}
	return prodv
}

func sum(v []int) int {
	prodv := 0
	for _, val := range v {
		prodv += val
	}
	return prodv
}

func placeWeights(weights, g1, g2, g3 []int, goal int) (int, int) {

	if len(weights) == 0 {
		evals.Inc()

		return prod(g3), len(g3)
	}

	bestv := math.MaxInt
	bestl := math.MaxInt

	if sum(g1)+weights[0] <= goal {
		resp1, len1 := placeWeights(weights[1:], append(g1, weights[0]), g2, g3, goal)
		if len1 < bestl || (len1 == bestl && resp1 < bestv) {
			bestv = resp1
			bestl = len1
		}
	}

	if sum(g2)+weights[0] <= goal && len(g1) > len(g2) {
		resp2, len2 := placeWeights(weights[1:], g1, append(g2, weights[0]), g3, goal)
		if len2 < bestl || (len2 == bestl && resp2 < bestv) {
			bestv = resp2
			bestl = len2
		}
	}

	if sum(g3)+weights[0] <= goal && len(g2) > len(g3) {
		resp3, len3 := placeWeights(weights[1:], g1, g2, append(g3, weights[0]), goal)
		if len3 < bestl || (len3 == bestl && resp3 < bestv) {
			bestv = resp3
			bestl = len3
		}
	}

	return bestv, bestl
}

func altGrouping(weights []int, goal int) int {
	i := 1
	for {
		evals.Set(float64(i))
		var built []int
		res := buildGrouping(weights, built, i, goal)
		if res > 0 {
			return res
		}
		i++
	}
}

func buildGrouping(weights, built []int, length, goal int) int {
	if len(built) == length {
		if sum(built) == goal {
			return prod(built)
		}
		return 0
	}

	best := math.MaxInt
	found := false
	for i, w := range weights {
		res := buildGrouping(weights[i+1:], append(built, w), length, goal)
		if res > 0 && res < best {
			best = res
			found = true
		}
	}

	if found {
		return best
	}
	return 0
}

func (s *Server) Solve2015day24part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-24.txt")
	if err != nil {
		return nil, err
	}

	var weights []int
	for _, line := range strings.Split(data, "\n") {
		val, _ := strconv.ParseInt(line, 10, 32)
		if val > 0 {
			weights = append(weights, int(val))
		}
	}

	s.CtxLog(ctx, fmt.Sprintf("LENGTH = %v", weights))

	return &pb.SolveResponse{BigAnswer: int64(computeGrouping(weights, 3))}, nil
}

func (s *Server) Solve2015day24part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-24.txt")
	if err != nil {
		return nil, err
	}

	var weights []int
	for _, line := range strings.Split(data, "\n") {
		val, _ := strconv.ParseInt(line, 10, 32)
		if val > 0 {
			weights = append(weights, int(val))
		}
	}

	s.CtxLog(ctx, fmt.Sprintf("LENGTH = %v", weights))

	return &pb.SolveResponse{BigAnswer: int64(computeGrouping(weights, 4))}, nil
}
