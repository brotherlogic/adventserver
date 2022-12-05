package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildCrates(data string) map[int][]string {
	crates := make(map[int][]string)

	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "1") {
			break
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		for i, char := range line {
			if char != ' ' && char != '[' && char != ']' {
				crateNum := (i - 1) / 4
				if _, ok := crates[crateNum]; !ok {
					crates[crateNum] = make([]string, 0)
				}
				crates[crateNum] = append(crates[crateNum], string(char))
			}
		}
	}

	return crates
}

func rearrangeCrates(data string, rev bool) string {
	crates := buildCrates(data)

	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, "move") {
			elems := strings.Fields(line)
			count, _ := strconv.ParseInt(elems[1], 10, 32)
			start, _ := strconv.ParseInt(elems[3], 10, 32)
			end, _ := strconv.ParseInt(elems[5], 10, 32)

			for i := 0; i < int(count); i++ {
				crate := crates[int(start)-1][0]
				crates[int(start)-1] = crates[int(start)-1][1:]
				crates[int(end)-1] = append([]string{crate}, crates[int(end)-1]...)
			}
		}
	}

	retstr := ""
	for i := 0; i < len(crates); i++ {
		retstr += crates[i][0]
	}

	return retstr
}

func (s *Server) Solve2022day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-5.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: rearrangeCrates(data, false)}, nil
}

func (s *Server) Solve2022day5part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-5.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: rearrangeCrates(data, true)}, nil
}
