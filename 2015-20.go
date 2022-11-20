package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	thouses = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day20_houses",
		Help: "The number of server requests",
	})
	tlen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day20_len",
		Help: "The number of server requests",
	})
)

func findMaxHouse(sval int) int {
	houses := make([]int, sval)

	for i := 1; i < len(houses); i++ {
		for j := i; j < len(houses); j += i {
			houses[j] += i * 10
		}

		if houses[i] >= sval {
			return i
		}
	}

	return -1
}

func findMaxElves(sval int) int {
	houses := make([]int, sval)

	for i := 1; i < len(houses); i++ {
		count := 0
		for j := i; j < len(houses); j += i {
			houses[j] += i * 10
			count++
			if count >= 50 {
				break
			}
		}

		if houses[i] >= sval {
			return i
		}
	}

	return -1
}

func (s *Server) Solve2015day20part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(findMaxHouse(36000000))}, nil
}

func (s *Server) Solve2015day20part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(findMaxElves(36000000))}, nil
}
