package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	nline = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2016_day23_line",
		Help: "The number of server requests",
	})
)

type toggler struct {
	a, b, c, d int
	program    []string
	cprogram   []string
	pointer    int
	output     string
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

func runToggleProgram(data string, init int) *toggler {
	toggler := &toggler{program: make([]string, 0), a: init}

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			toggler.program = append(toggler.program, strings.TrimSpace(line))
		}
	}

	toggler.cprogram = collapseProgram(toggler.program)

	for toggler.pointer < len(toggler.program) {
		nline.Set(float64(toggler.pointer))
		fields := strings.Fields(toggler.cprogram[toggler.pointer])
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
				if fields[2] == "a" {
					val2 = int64(toggler.a)
				}
				if fields[2] == "b" {
					val2 = int64(toggler.b)
				}
				if fields[2] == "c" {
					val2 = int64(toggler.c)
				}
				if fields[2] == "d" {
					val2 = int64(toggler.d)
				}
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
			jump := 0
			switch fields[1] {
			case "a":
				jump = toggler.a
			case "c":
				jump = toggler.c
			}

			if toggler.pointer+jump < len(toggler.program) {
				newline := toggler.program[toggler.pointer+jump]
				nfields := strings.Fields(newline)
				switch nfields[0] {
				case "inc":
					newline = strings.Replace(newline, "inc", "dec", -1)
				case "dec":
					newline = strings.Replace(newline, "dec", "inc", -1)
				case "tgl":
					newline = strings.Replace(newline, "tgl", "inc", -1)
				case "jnz":
					newline = strings.Replace(newline, "jnz", "cpy", -1)
				case "cpy":
					newline = strings.Replace(newline, "cpy", "jnz", -1)
				}

				toggler.program[toggler.pointer+jump] = newline
			}
			toggler.pointer++
			toggler.cprogram = collapseProgram(toggler.program)

		case "jnzcc":
			incrementer := fields[7]
			if fields[2] == incrementer {
				log.Fatalf("HUH")
			}
			val := toggler.get(fields[1])
			mult := toggler.get(fields[9])
			if fields[6] == "inc" {
				toggler.update(incrementer, val*mult)
			} else {
				toggler.update(incrementer, -val*mult)
			}
			toggler.pointer++
		case "jnzc":
			factor := fields[1]
			if fields[3] == factor {
				change := toggler.get(factor)
				if fields[4] == "inc" {
					toggler.update(fields[5], change)
				} else {
					toggler.update(fields[5], -change)
				}
			} else if fields[5] == factor {
				change := toggler.get(factor)
				if fields[2] == "inc" {
					toggler.update(fields[3], change)
				} else {
					toggler.update(fields[3], -change)
				}
			}
			toggler.pointer++

		default:
			log.Fatalf("Bad instruction: %v", fields[0])
		}
	}

	return toggler
}

func (t *toggler) get(register string) int {
	switch register {
	case "a":
		return t.a
	case "b":
		return t.b
	case "c":
		return t.c
	case "d":
		return t.d
	default:
		return getInt32(register)
	}
}

func (t *toggler) update(register string, val int) {
	switch register {
	case "a":
		t.a += val
	case "b":
		t.b += val
	case "c":
		t.c += val
	case "d":
		t.d += val
	default:
		log.Fatalf("Bad register on update: %v", register)
	}

}

func collapseProgram(program []string) []string {
	var nprog []string

	for i := range program {
		fields := strings.Fields(program[i])
		if strings.HasPrefix(program[i], "jnz") && strings.HasSuffix(program[i], "-2") {
			nprog = append(nprog, fmt.Sprintf("jnzc %v %v %v", fields[1], program[i-1], program[i-2]))
		} else {
			nprog = append(nprog, program[i])
		}
	}

	var nnprog []string
	for i := range nprog {
		if strings.HasPrefix(nprog[i], "jnz") && strings.HasSuffix(nprog[i], "-5") {
			fields := strings.Fields(program[i-1])
			nnprog = append(nnprog, fmt.Sprintf("jnzcc %v %v %v", fields[1], nprog[i-2], nprog[i-5]))
		} else {
			nnprog = append(nnprog, program[i])
		}
	}

	return nnprog
}

func (s *Server) Solve2016day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-23.txt")
	if err != nil {
		return nil, err
	}

	res := runToggleProgram(data, 7)
	return &pb.SolveResponse{Answer: int32(res.a)}, nil
}

func (s *Server) Solve2016day23part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-23.txt")
	if err != nil {
		return nil, err
	}

	res := runToggleProgram(data, 13)
	return &pb.SolveResponse{Answer: int32(res.a)}, nil
}
