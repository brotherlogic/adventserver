package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type elfProgram struct {
	a      int
	values []int
}

func (e elfProgram) getSignal() int {
	val := 0

	for _, sumv := range []int{20, 60, 100, 140, 180, 220} {
		val += sumv * e.values[sumv]
	}
	return val
}

func (e elfProgram) getScreen() string {
	screen := ""
	for i, val := range e.values[1:] {
		bound0 := i%40 - 1
		bound1 := i%40 + 1
		if i%40 == 0 {
			bound0 = 0
		}
		if i%40 == 39 {
			bound1 = 39
		}

		if val >= bound0 && val <= bound1 {
			screen += "#"
		} else {
			screen += "."
		}
		if i%40 == 39 {
			screen += "\n"
		}
	}

	return screen[:len(screen)-1]
}

func runElfProgram(data string) elfProgram {

	prog := elfProgram{a: 1, values: []int{1}}
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			fields := strings.Fields(line)
			switch fields[0] {
			case "addx":
				prog.values = append(prog.values, prog.a)
				prog.values = append(prog.values, prog.a)
				val, _ := strconv.ParseInt(fields[1], 10, 32)
				prog.a += int(val)
			case "noop":
				prog.values = append(prog.values, prog.a)
			}

		}
	}

	return prog
}

func (s *Server) Solve2022day10part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-10.txt")
	if err != nil {
		return nil, err
	}

	result := runElfProgram(data)

	return &pb.SolveResponse{Answer: int32(result.getSignal())}, nil
}

func (s *Server) Solve2022day10part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-10.txt")
	if err != nil {
		return nil, err
	}

	result := runElfProgram(data)

	return &pb.SolveResponse{StringAnswer: result.getScreen()}, nil
}
