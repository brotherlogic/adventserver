package main

import (
	"log"

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

func findMaxHouse(sval int32) int {
	houses := make(map[int]int32)

	start := 1
	for {
		thouses.Set(float64(start))
		tlen.Set(float64(len(houses)))
		for i := 1; i <= 10; i++ {
			houses[start*i] += int32(start)
		}

		if houses[start] >= sval {
			log.Printf("FOUND AT %v", houses[start])
			return start
		}
		delete(houses, start)

		start++
	}
}

func (s *Server) Solve2015day20part1(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{Answer: int32(findMaxHouse(36000000 / 10))}, nil
}
