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
	head := &elfNode{num: 1}
	h := head
	for i := 2; i <= num; i++ {
		next := &elfNode{num: i, prev: head}
		head.next = next
		head = next
	}

	head.next = h
	head.next.prev = head
	head = h
	count := num

	for {
		celf.Set(float64(count))
		if head.next.num == head.num {
			return head.num
		}
		jump := count / 2
		c := head
		for i := 0; i < jump; i++ {
			c = c.next
		}
		c.next.prev = c.prev
		c.prev.next = c.next

		count--
		head = head.next
	}
}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
func (s *Server) Solve2016day19part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runCircularPresents(3018458))}, nil
}
