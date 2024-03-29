package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	lengths = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day10_lengths",
		Help: "The number of server requests",
	})
)

func lookAndSay(s string) string {
	currRune := ' '
	currCount := 0

	var retStr strings.Builder
	for _, r := range s {
		if r == currRune {
			currCount++
		} else {
			if currCount > 0 {
				retStr.WriteString(fmt.Sprintf("%v", currCount))
				retStr.WriteRune(currRune)
			}
			currRune = r
			currCount = 1
		}
	}

	retStr.WriteString(fmt.Sprintf("%v", currCount))
	retStr.WriteRune(currRune)

	return retStr.String()
}

func (s *Server) Solve2015day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	start := "3113322113"
	for i := 0; i < 40; i++ {
		lengths.Set(float64(len(start)))
		start = lookAndSay(start)
	}

	return &pb.SolveResponse{BigAnswer: int64(len(start))}, nil
}

func (s *Server) Solve2015day10part2(ctx context.Context) (*pb.SolveResponse, error) {
	start := "3113322113"
	for i := 0; i < 50; i++ {
		lengths.Set(float64(len(start)))
		start = lookAndSay(start)
	}

	return &pb.SolveResponse{BigAnswer: int64(len(start))}, nil
}
