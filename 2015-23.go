package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	comCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day23_commands",
		Help: "The number of server requests",
	})
)

type computer struct {
	a, b int
}

func runProgram(ctx context.Context, program string, log func(context.Context, string), start int) computer {
	tcommands := strings.Split(program, "\n")
	var commands []string
	for i := range tcommands {
		nstr := strings.TrimSpace(tcommands[i])
		commands = append(commands, nstr)
	}

	curr := computer{a: start}
	i := 0
	for i < len(commands) {
		comCount.Inc()
		elems := strings.Fields(commands[i])
		if len(elems) == 0 {
			return curr
		}
		switch elems[0] {
		case "hlf":
			if elems[1] == "a" {
				curr.a = curr.a / 2
			} else {
				curr.b = curr.b / 2
			}
			i++
		case "tpl":
			if elems[1] == "a" {
				curr.a = curr.a * 3
			} else {
				curr.b = curr.b * 3
			}
			i++
		case "inc":
			if elems[1] == "a" {
				curr.a = curr.a + 1
			} else {
				curr.b = curr.b + 1
			}
			i++
		case "jmp":
			val, _ := strconv.ParseInt(elems[1], 10, 16)
			i += int(val)
		case "jio":
			if elems[1] == "a," && curr.a == 1 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else if elems[1] == "b," && curr.b == 1 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else {
				i++
			}
		case "jie":
			if elems[1] == "a," && curr.a%2 == 0 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else if elems[1] == "b," && curr.b%2 == 0 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else {
				i++
			}
		default:
			log(ctx, fmt.Sprintf("Unable to process %v", commands[i]))
		}
	}

	return curr
}

func (s *Server) Solve2015day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-23.txt")
	if err != nil {
		return nil, err
	}

	result := runProgram(ctx, data, s.CtxLog, 0)
	s.CtxLog(ctx, fmt.Sprintf("COMPUTED %+v", result))
	return &pb.SolveResponse{Answer: int32(result.b)}, nil
}

func (s *Server) Solve2015day23part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-23.txt")
	if err != nil {
		return nil, err
	}

	result := runProgram(ctx, data, s.CtxLog, 1)
	s.CtxLog(ctx, fmt.Sprintf("COMPUTED %+v", result))
	return &pb.SolveResponse{Answer: int32(result.b)}, nil
}
