package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"

	"golang.org/x/net/context"
)

func computeStringLength(str string) (int, int) {
	count := 0
	for r := 0; r < len(str); r++ {
		if str[r] == '\\' {
			if str[r+1] == 'x' {
				r += 3
			} else {
				r += 1
			}
		}
		count++
	}

	return len(str), count - 2
}

func computeStringAddition(str string) (int, int) {
	var nstr strings.Builder
	for i := 0; i < len(str); i++ {
		if str[i] == '"' {
			nstr.WriteString(`\"`)
		} else if str[i] == '\\' {
			nstr.WriteString(`\\`)
		} else {
			nstr.WriteString(string(str[i]))
		}
	}
	return len(str), len(nstr.String()) + 2
}

func (s *Server) Solve2015day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-8.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	chs := 0
	cds := 0
	for _, str := range strings.Split(trimmed, "\n") {
		c1, c2 := computeStringLength(str)
		chs += c1
		cds += c2
	}

	return &pb.SolveResponse{Answer: int32(chs - cds)}, nil
}

func (s *Server) Solve2015day8part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-8.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	chs := 0
	cds := 0
	for _, str := range strings.Split(trimmed, "\n") {
		c1, c2 := computeStringAddition(str)
		chs += c1
		cds += c2
	}

	return &pb.SolveResponse{Answer: int32(cds - chs)}, nil
}
