package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func in(val int, arr []int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}

	return false
}

func reduceTo(arr []string, word string) []string {
	var newa []string
	for _, char := range arr {
		if strings.Contains(word, string(char)) {
			newa = append(newa, string(char))
		}
	}

	return newa
}

func remove(arr []string, word string) []string {
	var newa []string
	for _, char := range arr {
		if !strings.Contains(word, string(char)) {
			newa = append(newa, string(char))
		}
	}

	return newa
}

func reduce(theory map[int][]string, ind []int, word string) {
	for key, arr := range theory {
		if in(key, ind) {
			theory[key] = reduceTo(arr, word)
		} else {
			theory[key] = remove(arr, word)
		}
	}
}

func buildCounts(data string) int {
	easyCount := 0
	for _, entry := range strings.Split(data, "\n") {
		theory := make(map[int][]string)
		for i := 0; i <= 6; i++ {
			theory[i] = []string{"a", "b", "c", "d", "e", "f", "g"}
		}
		pieces := strings.Split(entry, "|")
		for _, word := range strings.Fields((strings.TrimSpace(pieces[0]))) {
			switch len(word) {
			case 2:
				// This is 1
				reduce(theory, []int{2, 5}, word)
			case 3:
				//This is 7
				reduce(theory, []int{0, 2, 5}, word)
			case 4:
				//This is 4
				reduce(theory, []int{1, 2, 3, 5}, word)
			case 5:
				// This is 2, 3, 5

			case 6:
				// This is 0,6,9
			case 7:
				//This is eight - we learn nothing here
			}
		}
	}

	for _, entry := range strings.Split(data, "\n") {
		pieces := strings.Split(entry, "|")
		for _, word := range strings.Fields((strings.TrimSpace(pieces[1]))) {
			switch len(word) {
			case 2, 3, 4, 7:
				easyCount++
			}
		}
	}

	return easyCount
}

func (s *Server) Solve2021day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-8.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(buildCount(trimmed))}, nil
}
