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
)

func findMaxHouse(sval int) int {
	houses := make(map[int]int)

	start := 1
	for {
		thouses.Set(float64(start))
		for i := 1; i <= 10; i++ {
			houses[start*i] += start * 10
			if houses[start*i] >= sval {
				return start * i
			}
		}

		start++
	}
}

func (s *Server) Solve2015day20part1(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{Answer: int32(findMaxHouse(36000000))}, nil
}
