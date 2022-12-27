package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type registers struct {
	register map[string]int
}

func (r *registers) runLine(line string) {
	fields := strings.Fields(line)

	registerCheck := fields[4]
	comp := fields[5]
	value := getInt32(fields[6])

	apply := false

	switch comp {
	case ">":
		apply = r.register[registerCheck] > value
	case "<":
		apply = r.register[registerCheck] < value
	case ">=":
		apply = r.register[registerCheck] >= value
	case "<=":
		apply = r.register[registerCheck] <= value
	case "==":
		apply = r.register[registerCheck] == value
	case "!=":
		apply = r.register[registerCheck] != value
	}

	if apply {
		register := fields[0]
		value := getInt32(fields[2])
		if fields[1] == "inc" {
			r.register[register] += value
		} else {
			r.register[register] -= value
		}
	}

}

func runJumpProgram(data string) *registers {
	registers := &registers{register: make(map[string]int)}

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		registers.runLine(strings.TrimSpace(line))
	}
	return registers
}

func (s *Server) Solve2017day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-7.txt")
	if err != nil {
		return nil, err
	}
	res := runJumpProgram(data)

	highest := 0
	for _, value := range res.register {
		if value > highest {
			highest = value
		}
	}

	return &pb.SolveResponse{Answer: int32(highest)}, nil
}
