package main

import (
	"log"
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

func convertToMap(line string) map[string]int64 {
	mapper := make(map[string]int64)
	for i := 0; i < len(line)-1; i++ {
		mapper[string(line[i])+string(line[i+1])] = 1
	}
	return mapper
}

func runData(data string, count int) map[string]int64 {
	first := strings.Split(data, "\n")[0]
	mapper := convertToMap(strings.TrimSpace(first))
	rules := buildRules(data)

	for i := 0; i < count; i++ {
		mapper = runRules(mapper, rules)
	}

	return mapper
}

func runRules(mapper map[string]int64, rules map[string]string) map[string]int64 {
	nmap := make(map[string]int64)
	for val, count := range mapper {
		if new, ok := rules[val]; ok {
			nmap[string(val[0])+new] += count
			nmap[new+string(val[1])] += count
		}
	}
	return nmap
}

func getCommons(line map[string]int64) (int64, int64) {
	counts := make(map[string]int64)
	for str, co := range line {
		counts[string(str[0])] += co
		counts[string(str[1])] += co
	}

	highest := int64(0)
	lowest := int64(math.MaxInt64)

	for key, v := range counts {
		if v > highest {
			highest = v
			log.Printf("HIGH %v -> %v", key, (v+1)/2)
		}
		if v < lowest {
			lowest = v
			log.Printf("LOW %v -> %v", key, (v+1)/2)
		}
	}

	return (highest + 1) / 2, (lowest + 1) / 2
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

func (s *Server) Solve2021day14part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-14.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	newone := runData(trimmed, 40)
	mc, lc := getCommons(newone)
	return &pb.SolveResponse{Answer: int32(mc - lc)}, nil
}
