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
