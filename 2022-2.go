package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getRPSScore(data string) int {
	score := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			elems := strings.Fields(line)
			switch elems[0] {
			case "A":
				switch elems[1] {
				case "X":
					score += 1 + 3
				case "Y":
					score += 2 + 6
				case "Z":
					score += 3 + 0
				}
			case "B":
				switch elems[1] {
				case "X":
					score += 1 + 0
				case "Y":
					score += 2 + 3
				case "Z":
					score += 3 + 6
				}
			case "C":
				switch elems[1] {
				case "X":
					score += 1 + 6
				case "Y":
					score += 2 + 0
				case "Z":
					score += 3 + 3
				}
			}
		}
	}

	return score
}

func getRPSScoreAmend(data string) int {
	score := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			elems := strings.Fields(line)
			switch elems[0] {
			case "A":
				switch elems[1] {
				case "X":
					score += 3 + 0
				case "Y":
					score += 1 + 3
				case "Z":
					score += 2 + 6
				}
			case "B":
				switch elems[1] {
				case "X":
					score += 1 + 0
				case "Y":
					score += 2 + 3
				case "Z":
					score += 3 + 6
				}
			case "C":
				switch elems[1] {
				case "X":
					score += 2 + 0
				case "Y":
					score += 3 + 3
				case "Z":
					score += 1 + 6
				}
			}
		}
	}

	return score
}

func (s *Server) Solve2022day2part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-2.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getRPSScore(data))}, nil
}

func (s *Server) Solve2022day2part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-2.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getRPSScoreAmend(data))}, nil
}
