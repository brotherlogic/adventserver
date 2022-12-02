package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type disc struct {
	positions int
	start     int
}

func (d disc) fall(t int) bool {
	return (d.start+t)%d.positions == 0
}

func runFall(t int, discs []disc) bool {
	for i, disc := range discs {
		if !disc.fall(t + i + 1) {
			return false
		}
	}

	return true
}

func fallBalls(data string) int {
	var discs []disc

	for _, line := range strings.Split(data, "\n") {
		if len(line) > 0 {
			elems := strings.Fields(line)
			pos, _ := strconv.ParseInt(elems[3], 10, 32)
			p, _ := strconv.ParseInt(strings.ReplaceAll(elems[11], ".", ""), 10, 32)
			discs = append(discs, disc{positions: int(pos), start: int(p)})
		}
	}

	t := 0
	for {
		if runFall(t, discs) {
			return t
		}
		t++
	}
}

func (s *Server) Solve2016day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-15.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(fallBalls(data))}, nil
}
