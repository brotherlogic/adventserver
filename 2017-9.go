package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func convertStream(line string) int {
	inGarbage := false

	groups := 0
	depth := 0
	pointer := 0
	for pointer < len(line) {
		switch line[pointer] {
		case '!':
			pointer++
		case '{':
			if !inGarbage {
				depth++
			}
		case '}':
			if !inGarbage {
				groups += depth
				depth--
			}
		case '<':
			inGarbage = true
		case '>':
			inGarbage = false
		}

		pointer++
	}

	return 0
}

func (s *Server) Solve2017day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-9.txt")
	if err != nil {
		return nil, err
	}
	sumv := int32(0)
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		sumv += int32(convertStream(strings.TrimSpace(line)))
	}
	return &pb.SolveResponse{Answer: sumv}, nil
}
