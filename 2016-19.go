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

type elfNode struct {
	num  int
	next *elfNode
	prev *elfNode
}

func runCircularPresents(num int) int {
	count := num
	head := &elfNode{num: 1}
	c := head
	for i := 2; i <= num; i++ {
		newOne := &elfNode{num: i, prev: c}

		c.next = newOne
		c = newOne
	}
	head.prev = c
	c.next = head

	start := head
	for i := 0; i < num/2; i++ {
		start = start.next
	}

	if num%2 == 0 {
		start = start.prev
	}

	for start.next.num != start.num {
		celf.Set(float64(count))
		remove := start
		if count%2 == 0 {
			remove = remove.next
		}

		remove.prev.next = remove.next
		remove.next.prev = remove.prev
		start = remove.next
		count--
	}

	return start.num
}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
func (s *Server) Solve2016day19part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runCircularPresents(3018458))}, nil
}
