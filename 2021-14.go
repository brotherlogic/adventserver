package main

import (
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildRules(data string) map[string]string {
	rules := make(map[string]string)

	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "->") {
			elems := strings.Split(line, "->")
			rules[strings.TrimSpace(elems[0])] = strings.TrimSpace(elems[1])
		}
	}
	return rules
}

func runData(data string, count int) string {
	first := strings.Split(data, "\n")[0]
	rules := buildRules(data)

	for i := 0; i < count; i++ {
		first = runRules(first, rules)
	}

	return first
}

func runRules(line string, rules map[string]string) string {
	nStr := string(line[0])

	for i := 0; i < len(line)-1; i++ {
		mapper := line[i : i+2]
		if newChar, ok := rules[mapper]; ok {
			nStr += newChar + string(line[i+1])
		}
	}

	return nStr
}

func getCommons(line string) (int, int) {
	counts := make(map[string]int)
	for _, ch := range line {
		counts[string(ch)]++
	}

	highest := 0
	lowest := math.MaxInt32

	for _, v := range counts {
		if v > highest {
			highest = v
		}
		if v < lowest {
			lowest = v
		}
	}

	return highest, lowest
}

func (s *Server) Solve2021day14part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-14.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	newone := runData(trimmed, 10)
	mc, lc := getCommons(newone)
	return &pb.SolveResponse{Answer: int32(mc - lc)}, nil
}
