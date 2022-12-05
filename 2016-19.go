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
	nextElf := make(map[int]int)
	prevElf := make(map[int]int)
	for i := 0; i < num; i++ {
		nextElf[i] = (i + 1) % num
		prevElf[(i+1)%num] = i
	}

	currElf := 0
	for len(nextElf) > 1 {
		jump := len(nextElf) / 2
		find := currElf
		for i := 0; i < jump; i++ {
			find = nextElf[find]
		}

		nextElf[prevElf[find]] = nextElf[find]
		prevElf[nextElf[find]] = prevElf[find]
		delete(nextElf, find)
		delete(prevElf, find)

		currElf = nextElf[currElf]
	}

	for num := range nextElf {
		return num + 1
	}

	return -1
}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
func (s *Server) Solve2016day19part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runCircularPresents(3018458))}, nil
}
