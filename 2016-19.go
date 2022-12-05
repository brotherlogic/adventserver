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
	elves := make([]bool, num)
	pointer := 0
	find := 0

	for {
		celf.Set(float64(num - find))
		if find == num-1 {
			for i, elf := range elves {
				if !elf {
					return i + 1
				}
			}
		}

		for {
			// Find the next in play elf
			if elves[pointer] {
				pointer++
				pointer = pointer % num
			} else {
				break
			}
		}

		// Get the elf to steal from
		stealJump := (num - find) / 2
		sPointer := pointer
		for {
			if elves[sPointer] {
				sPointer++
				sPointer = sPointer % num
			} else {
				if stealJump == 0 {
					break
				} else {
					stealJump--
					sPointer++
					sPointer = sPointer % num
				}
			}
		}

		find++
		elves[sPointer] = true

		pointer++
		pointer = pointer % num
	}
}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
func (s *Server) Solve2016day19part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runCircularPresents(3018458))}, nil
}
