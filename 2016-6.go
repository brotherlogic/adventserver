package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getMostCommon(data []string, pos int) string {
	counts := make(map[string]int)

	for _, dat := range data {
		counts[string(dat[pos])]++
	}

	best := 0
	bounce := ""
	for key, count := range counts {
		if count > best {
			best = count
			bounce = key
		}
	}

	return bounce
}

func getCommon(data string) string {
	elems := strings.Split(data, "\n")

	answer := ""
	for i := 0; i < len(elems[0]); i++ {
		answer += getMostCommon(elems, i)
	}

	return answer
}
func (s *Server) Solve2016day6part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-6.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: getCommon(data)}, nil
}
