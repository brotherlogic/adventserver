package main

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	line = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day12_line",
		Help: "The number of server requests",
	})
	evals12 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day12_evals",
		Help: "The number of server requests",
	})
)

type mstate struct {
	a, b, c, d int
}

func (m *mstate) set(reg string, value int) {
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

func runMonorailProgram(data string, set bool) *mstate {

	ppoint := 0
	lines := strings.Split(strings.TrimSpace(data), "\n")
	mstate := &mstate{}
	if set {
		mstate.c = 1
	}

	for ppoint < len(lines) {
		evals12.Inc()
		line.Set(float64(ppoint))
		fields := strings.Fields(lines[ppoint])
		switch fields[0] {
		case "jnz":
			val, _ := strconv.ParseInt(fields[2], 10, 32)
			switch fields[1] {
			case "a":
				if mstate.a != 0 {
					ppoint += int(val)
				} else {
					ppoint++
				}
			case "b":
				if mstate.b != 0 {
					ppoint += int(val)
				} else {
					ppoint++
				}
			case "c":
				if mstate.c != 0 {
					ppoint += int(val)
				} else {
					ppoint++
				}
			case "d":
				if mstate.d != 0 {
					ppoint += int(val)
				} else {
					ppoint++
				}
			default:
				val, _ := strconv.ParseInt(fields[1], 10, 32)
				val2, _ := strconv.ParseInt(fields[2], 10, 32)
				if val != 0 {
					ppoint += int(val2)
				} else {
					ppoint++
				}
			}
		case "cpy":
			if fields[1] == "a" {
				mstate.set(fields[2], mstate.a)
			} else if fields[1] == "b" {
				mstate.set(fields[2], mstate.b)
			} else if fields[1] == "c" {
				mstate.set(fields[2], mstate.c)
			} else if fields[1] == "d" {
				mstate.set(fields[2], mstate.d)
			} else {
				val, _ := strconv.ParseInt(fields[1], 10, 32)
				mstate.set(fields[2], int(val))
			}
			ppoint++
		case "inc":
			switch fields[1] {
			case "a":
				mstate.a++
			case "b":
				mstate.b++
			case "c":
				mstate.c++
			case "d":
				mstate.d++
			}
			ppoint++
		case "dec":
			switch fields[1] {
			case "a":
				mstate.a--
			case "b":
				mstate.b--
			case "c":
				mstate.c--
			case "d":
				mstate.d--
			}
			ppoint++
		}
	}

	return mstate
}

func (s *Server) Solve2016day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-12.txt")
	if err != nil {
		return nil, err
	}

	state := runMonorailProgram(data, false)

	return &pb.SolveResponse{Answer: int32(state.a)}, nil
}

func (s *Server) Solve2016day12part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-12.txt")
	if err != nil {
		return nil, err
	}

	state := runMonorailProgram(data, true)

	return &pb.SolveResponse{Answer: int32(state.a)}, nil
}
