package main

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type toggler struct {
	a, b, c, d int
	program    []string
	pointer    int
}

func (m *toggler) set(reg string, value int) {
	switch reg {
	case "a":
		m.a = value
	case "b":
		m.b = value
	case "c":
		m.c = value
	case "d":
		m.d = value
	default:
		log.Fatalf("Bad set: %v", reg)
	}
}

func runToggleProgram(data string) *toggler {
	toggler := &toggler{program: make([]string, 0)}

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			toggler.program = append(toggler.program, strings.TrimSpace(line))
		}
	}

	for toggler.pointer < len(toggler.program) {
		fields := strings.Fields(toggler.program[toggler.pointer])
		switch fields[0] {
		case "jnz":
			val, _ := strconv.ParseInt(fields[2], 10, 32)
			switch fields[1] {
			case "a":
				if toggler.a != 0 {
					toggler.pointer += int(val)
				} else {
					toggler.pointer++
				}
			case "b":
				if toggler.b != 0 {
					toggler.pointer += int(val)
				} else {
					toggler.pointer++
				}
			case "c":
				if toggler.c != 0 {
					toggler.pointer += int(val)
				} else {
					toggler.pointer++
				}
			case "d":
				if toggler.d != 0 {
					toggler.pointer += int(val)
				} else {
					toggler.pointer++
				}
			default:
				val, _ := strconv.ParseInt(fields[1], 10, 32)
				val2, _ := strconv.ParseInt(fields[2], 10, 32)
				if val != 0 {
					toggler.pointer += int(val2)
				} else {
					toggler.pointer++
				}
			}
		case "cpy":
			if fields[1] == "a" {
				toggler.set(fields[2], toggler.a)
			} else if fields[1] == "b" {
				toggler.set(fields[2], toggler.b)
			} else if fields[1] == "c" {
				toggler.set(fields[2], toggler.c)
			} else if fields[1] == "d" {
				toggler.set(fields[2], toggler.d)
			} else {
				val, _ := strconv.ParseInt(fields[1], 10, 32)
				toggler.set(fields[2], int(val))
			}
			toggler.pointer++
		case "inc":
			switch fields[1] {
			case "a":
				toggler.a++
			case "b":
				toggler.b++
			case "c":
				toggler.c++
			case "d":
				toggler.d++
			}
			toggler.pointer++
		case "dec":
			switch fields[1] {
			case "a":
				toggler.a--
			case "b":
				toggler.b--
			case "c":
				toggler.c--
			case "d":
				toggler.d--
			}
			toggler.pointer++
		case "tgl":
			toggler.pointer++
		}
	}

	return toggler
}

func (s *Server) Solve2016day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-23.txt")
	if err != nil {
		return nil, err
	}

	res := runToggleProgram(data)
	return &pb.SolveResponse{Answer: int32(res.a)}, nil
}
