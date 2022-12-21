package main

import (
	"fmt"
	"log"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func unencrpyt(data string) int {
	var arr []int

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			arr = append(arr, getInt32(line))
		}
	}

	numMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = i
	}

	for _, val := range arr {
		log.Printf("MOVE %v", printNumMap(numMap))
		numMap = moveMap(numMap, val)
	}
	log.Printf("DONE %v", printNumMap(numMap))

	narr := make([]int, len(arr))
	for key, val := range numMap {
		narr[val] = key
	}

	log.Printf("FOUND %v", narr)

	return narr[(numMap[0]+1000)%len(narr)] + narr[(numMap[0]+2000)%len(narr)] + narr[(numMap[0]+3000)%len(narr)]
}

func printNumMap(numMap map[int]int) string {
	narr := make([]int, len(numMap))
	for key, val := range numMap {
		narr[val] = key
	}
	return fmt.Sprintf("%v", narr)
}

func moveNumber(arr []int, num int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = i
	}

	numMap = moveMap(numMap, num)

	newarr := make([]int, len(arr))
	for key, val := range numMap {
		newarr[val] = key
	}

	return newarr
}

func moveMap(numMap map[int]int, num int) map[int]int {
	if num == 0 {
		return numMap
	}

	nIndex := numMap[num]
	llen := len(numMap)

	newIndex := (nIndex + num) % llen
	if num < 0 {
		newIndex = (nIndex + num - 1) % llen
	}
	if newIndex < 0 {
		newIndex = llen + newIndex
	}

	log.Printf("NEW %v (%v) -> %v", num, nIndex, newIndex)
	following := 0
	for key, val := range numMap {
		if val == newIndex {
			following = key
		}
	}

	log.Printf("INSERT %v-> %v", num, following)

	// Remove the number
	delete(numMap, num)
	for key, val := range numMap {
		if val > nIndex {
			numMap[key] = val - 1
		}
	}

	newNewIndex := numMap[following] + 1

	for key, val := range numMap {
		if val >= newNewIndex {
			numMap[key] = val + 1
		}
	}
	numMap[num] = newNewIndex

	return numMap
}

func (s *Server) Solve2022day20part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-20.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(unencrpyt(data))}, nil
}
