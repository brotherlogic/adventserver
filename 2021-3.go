package main

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func count(lines []string, pos int) map[string]int {
	mapper := make(map[string]int)
	for _, line := range lines {
		mapper[string(line[pos])]++
	}
	return mapper
}

func mostCommon(lines []string, pos int) string {
	counts := count(lines, pos)
	if counts["0"] > counts["1"] {
		return "0"
	}
	return "1"
}

func leastCommon(lines []string, pos int) string {
	counts := count(lines, pos)
	if counts["0"] <= counts["1"] {
		return "0"
	}
	return "1"
}

func convBinary(bstr string) int {
	val, _ := strconv.ParseInt(bstr, 2, 64)
	return int(val)
}

func computePower(data string) int {
	lines := strings.Split(data, "\n")

	ms := ""
	ls := ""

	for pos := 0; pos < len(lines[0]); pos++ {
		ms += mostCommon(lines, pos)
		ls += leastCommon(lines, pos)
	}

	return convBinary(ms) * convBinary(ls)
}

func computeOGRSub(data []string, pos int) string {
	log.Printf("%v", len(data))
	if len(data) == 1 {
		return data[0]
	}

	mc := mostCommon(data, pos)
	var ndata []string
	for _, d := range data {
		if string(d[pos]) == mc {
			ndata = append(ndata, d)
		}
	}

	return computeOGRSub(ndata, pos+1)
}

func computeOGR(data string) int {
	return convBinary(computeOGRSub(strings.Split(data, "\n"), 0))
}

func computeLSRSub(data []string, pos int) string {
	if len(data) == 1 {
		return data[0]
	}

	mc := leastCommon(data, pos)
	var ndata []string
	for _, d := range data {
		if string(d[pos]) == mc {
			ndata = append(ndata, d)
		}
	}

	return computeLSRSub(ndata, pos+1)
}

func computeLSR(data string) int {
	return convBinary(computeLSRSub(strings.Split(data, "\n"), 0))
}

func (s *Server) Solve2021day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-3.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computePower(trimmed))}, nil
}

func (s *Server) Solve2021day3part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-3.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeLSR(trimmed) * computeOGR(trimmed))}, nil
}
