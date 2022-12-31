package main

import (
	"log"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func reverseArr(arr []int, start, end int) []int {
	rcount := end - start
	if rcount < 0 {
		rcount = rcount + len(arr)
	}

	for i := 0; i < rcount; i += 2 {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--

		start = start % len(arr)
		if end < 0 {
			end = len(arr) + end
		}
	}
	return arr
}

func runArray(list, steps []int) int32 {
	cpointer := 0

	for i, val := range steps {
		log.Printf("%v -> %v-%v", val, cpointer, (cpointer+val)%len(list)-1)
		if val != 0 {
			list = reverseArr(list, cpointer, (cpointer+val)%len(list)-1)
		}
		cpointer += val + i
		cpointer = cpointer % (len(list))
	}

	return int32(list[0] * list[1])
}

func (s *Server) Solve2017day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	var list []int
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	return &pb.SolveResponse{Answer: runArray(list, []int{212, 254, 178, 237, 2, 0, 1, 54, 167, 92, 117, 125, 255, 61, 159, 164})}, nil
}
