package main

import (
	"sort"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func sortWord(str string) string {
	var intarr []int
	for _, c := range str {
		intarr = append(intarr, int(c))
	}

	sort.Ints(intarr)

	strret := ""
	for _, elem := range intarr {
		strret += string(rune(elem))
	}
	return strret
}

func isValidPassword(in string, srt bool) bool {
	counts := make(map[string]int)
	for _, word := range strings.Fields(in) {
		if srt {
			word = sortWord(word)
		}
		counts[word]++
	}

	for _, val := range counts {
		if val > 1 {
			return false
		}
	}

	return true
}

func (s *Server) Solve2017day4part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-d4.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	count := 0
	for _, str := range strings.Split(trimmed, "\n") {
		if isValidPassword(str, false) {
			count++
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}

func (s *Server) Solve2017day4part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-d4.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	count := 0
	for _, str := range strings.Split(trimmed, "\n") {
		if isValidPassword(str, true) {
			count++
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
