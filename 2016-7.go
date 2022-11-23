package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func sslSupport(in string) bool {
	inBracks := false
	bracks := make(map[string]string)
	iBracks := make(map[string]bool)
	for i := 0; i < len(in)-2; i++ {
		if in[i] == '[' {
			inBracks = true
		} else if in[i] == ']' {
			inBracks = false
		} else if in[i] == in[i+2] && in[i] != in[i+1] {
			if inBracks {
				iBracks[in[i:i+3]] = true
			} else {
				bracks[in[i:i+3]] = fmt.Sprintf("%v%v%v", string(in[i+1]), string(in[i]), string(in[i+1]))
			}
		}
	}

	for _, end := range bracks {
		if _, ok := iBracks[end]; ok {
			return true
		}
	}

	return false
}

func tlsSupport(in string) bool {
	inBracks := false
	found := false
	for i := 0; i < len(in)-3; i++ {
		if in[i] == '[' {
			inBracks = true
		} else if in[i] == ']' {
			inBracks = false
		} else if in[i] == in[i+3] && in[i+1] == in[i+2] && in[i] != in[i+1] {
			found = true
			if inBracks {
				return false
			}
		}
	}
	return found
}

func (s *Server) Solve2016day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-7.txt")
	if err != nil {
		return nil, err
	}

	count := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			if tlsSupport(line) {
				s.CtxLog(ctx, fmt.Sprintf("%v", line))
				count++
			}
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}

func (s *Server) Solve2016day7part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-7.txt")
	if err != nil {
		return nil, err
	}

	count := 0
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			if sslSupport(line) {
				s.CtxLog(ctx, fmt.Sprintf("%v", line))
				count++
			}
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
