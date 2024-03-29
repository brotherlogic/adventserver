package main

import (
	"log"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countOn(array [][]bool) int {
	count := 0
	for _, val := range array {
		for _, nval := range val {
			if nval {
				count++
			}
		}
	}
	return count
}

func buildArray(data string, stuck bool) [][]bool {
	len1 := len(strings.Split(data, "\n")[0])

	var arr [][]bool
	for i := 0; i < len1+2; i++ {
		var narr []bool
		for j := 0; j < len1+2; j++ {
			narr = append(narr, false)
		}
		arr = append(arr, narr)
	}

	for i, val := range strings.Split(data, "\n") {
		for j, ch := range strings.TrimSpace(val) {
			if ch == '#' {
				arr[i+1][j+1] = true
			} else if ch != '.' {
				log.Fatalf("Nope")
			}
		}
	}

	if stuck {
		arr[1][1] = true
		arr[1][len1] = true
		arr[len1][1] = true
		arr[len1][len1] = true
	}

	return arr
}

func getValue(arr [][]bool, i, j int, stuck bool) bool {
	if i == 0 || j == 0 || i == len(arr)-1 || j == len(arr)-1 {
		return false
	}

	if stuck {
		if (i == 1 && j == 1) ||
			(i == 1 && j == len(arr)-2) ||
			(i == len(arr)-2 && j == 1) ||
			(i == len(arr)-2 && j == len(arr)-2) {
			return true
		}
	}

	count := 0
	for ic := i - 1; ic <= i+1; ic++ {
		for jc := j - 1; jc <= j+1; jc++ {
			if ic != i || jc != j {
				if arr[ic][jc] {
					count++
				}
			}
		}
	}

	if arr[i][j] && (count == 2 || count == 3) {
		return true
	}

	if !arr[i][j] && count == 3 {
		return true
	}

	return false
}

func rotateArray(arr [][]bool, stuck bool) [][]bool {
	var nar [][]bool
	for i := 0; i < len(arr); i++ {
		var narr []bool
		for j := 0; j < len(arr); j++ {
			narr = append(narr, getValue(arr, i, j, stuck))
		}
		nar = append(nar, narr)
	}

	return nar
}

func rotate(data string, times int, stuck bool) int {

	array := buildArray(data, stuck)

	for i := 0; i < times; i++ {
		array = rotateArray(array, stuck)
	}

	return countOn(array)
}

func (s *Server) Solve2015day18part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-18.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(rotate(trimmed, 100, false))}, nil
}

func (s *Server) Solve2015day18part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-18.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(rotate(trimmed, 100, true))}, nil
}
