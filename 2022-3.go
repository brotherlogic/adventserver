package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func sumOfPriorities(data string) int {
	sumv := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			sumv += getPriority(getPCommon(line))
		}
	}

	return sumv
}

func getPCommon(line string) string {
	halfOne := line[:len(line)/2]
	halfTwo := line[len(line)/2:]
	for _, char := range halfOne {
		if strings.Contains(halfTwo, string(char)) {
			return string(char)
		}
	}
	return ""
}

func sumOfCommons(data string) int {
	elems := strings.Split(data, "\n")
	sumv := 0
	for i := 0; i < len(elems); i += 3 {
		sumv += getPriority(getFCommon(elems[i], elems[i+1], elems[i+2]))
	}
	return sumv
}

func getFCommon(line1, line2, line3 string) string {
	found := make(map[rune]int)
	for _, char := range strings.TrimSpace(line1) {
		found[char] = 1
	}
	for _, char := range strings.TrimSpace(line2) {
		if found[char] == 1 {
			found[char]++
		}
	}
	for _, char := range strings.TrimSpace(line3) {
		if found[char] == 2 {
			found[char]++
		}
	}
	for k, v := range found {
		if v == 3 {
			return string(k)
		}
	}
	return ""
}

func getPriority(char string) int {
	rune := int(char[0])
	if (rune) <= 'z' && (rune) >= 'a' {
		return rune - 96
	} else {
		return rune - 38
	}
}

func (s *Server) Solve2022day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-3.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(sumOfPriorities(data))}, nil
}
