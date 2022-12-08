package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func translateCode(code, trans string) string {
	return code
}

func fullTranslate(data, code string) string {
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			code = translateCode(code, strings.TrimSpace(line))
		}
	}
	return code
}

func (s *Server) Solve2016day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-21.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: fullTranslate(data, "abcdefgh")}, nil
}
