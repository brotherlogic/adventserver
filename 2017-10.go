package main

import (
	"strconv"

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
		if val != 0 {
			list = reverseArr(list, cpointer, (cpointer+val-1)%len(list))
		}
		cpointer += val + i
		cpointer = cpointer % (len(list))
	}

	return int32(list[0] * list[1])
}

func runHash(in string, steps []int) string {
	var barr []int
	for _, char := range in {
		barr = append(barr, int(char))
	}
	barr = append(barr, steps...)

	var list []int
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	step := 0
	cpointer := 0
	for i := 0; i < 64; i++ {
		for _, val := range barr {
			if val != 0 {
				list = reverseArr(list, cpointer, (cpointer+val-1)%len(list))
			}
			cpointer += val + step
			cpointer = cpointer % len(list)
			step++
		}
	}

	var newarr []int64
	for i := 0; i < len(list); i += 16 {
		val := list[i]
		for j := 1; j < 16; j++ {
			val ^= list[i+j]
		}

		newarr = append(newarr, int64(val))
	}

	ret := ""
	for _, num := range newarr {
		if num < 16 {
			ret += "0" + strconv.FormatInt(num, 16)
		} else {
			ret += strconv.FormatInt(num, 16)
		}
	}

	return ret
}

func (s *Server) Solve2017day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	var list []int
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	return &pb.SolveResponse{Answer: runArray(list, []int{212, 254, 178, 237, 2, 0, 1, 54, 167, 92, 117, 125, 255, 61, 159, 164})}, nil
}
func (s *Server) Solve2017day10part2(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{StringAnswer: runHash("212,254,178,237,2,0,1,54,167,92,117,125,255,61,159,164", []int{17, 31, 73, 47, 23})}, nil
}
