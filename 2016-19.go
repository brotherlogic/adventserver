package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	celf = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2016_day19_current",
		Help: "The number of server requests",
	})
)

func runPresents(num int) int {
	elves := make([]bool, num)
	count := 0
	pointer := 0
	keep := true

	for {
		for {
			// Find the next in play elf
			if elves[pointer] {
				pointer++
				pointer = pointer % num
			} else {
				break
			}
		}

		if keep {
			keep = false
		} else {
			keep = true
			elves[pointer] = true
			count++
			if count == num-1 {
				for i, val := range elves {
					if !val {
						return i + 1
					}
				}
			}
		}
		pointer++
		pointer = pointer % num
	}
}

func runCircularPresents(num int) int {
	elves := make([]int, num)
	for i := 1; i <= num; i++ {
		elves[i-1] = i
	}
	pointer := 0

	for len(elves) > 1 {
		celf.Set(float64(len(elves)))
		toRemove := (pointer + len(elves)/2) % len(elves)
		pointer++
		pointer = pointer % len(elves)
		elves = append(elves[0:toRemove], elves[toRemove+1:]...)

	}

	return elves[0]
}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
func (s *Server) Solve2016day19part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runCircularPresents(3018458))}, nil
}
