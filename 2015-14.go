package main

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
)

type reindeer struct {
	speed     int
	blastTime int
	waitTime  int
	name      string
}

func getRein(line string) *reindeer {
	elems := strings.Fields(line)
	spd, _ := strconv.ParseInt(elems[3], 10, 32)
	blst, _ := strconv.ParseInt(elems[6], 10, 32)
	wait, _ := strconv.ParseInt(elems[13], 10, 32)

	return &reindeer{speed: int(spd), blastTime: int(blst), waitTime: int(wait), name: elems[0]}
}

func computeDistance(r *reindeer, time int) int {
	segments := time / (r.blastTime + r.waitTime)
	remainder := time % (r.blastTime + r.waitTime)

	little := 0
	if remainder < r.blastTime {
		little = (r.speed * remainder)
	} else {
		little = r.speed * r.blastTime
	}

	return little + segments*(r.speed*r.blastTime)
}

func getDistance(data string, time int) int32 {
	bd := 0
	for _, line := range strings.Split(data, "\n") {
		b := computeDistance(getRein(line), time)
		if b > bd {
			bd = b
		}
	}

	return int32(bd)
}

func getPoints(data string, time int) []string {
	var winners []string

	best := 0
	for _, line := range strings.Split(data, "\n") {
		rein := getRein(line)
		b := computeDistance(rein, time)
		if b > best {
			best = b
			winners = []string{rein.name}
		} else if b == best {
			winners = append(winners, rein.name)
		}
	}

	return winners
}

func runPoints(data string, mtime int) int32 {
	winners := make(map[string]int)
	for t := 1; t <= mtime; t++ {
		for _, w := range getPoints(data, t) {
			winners[w]++
		}
	}

	bp := 0
	for _, points := range winners {
		if points > bp {
			bp = points
		}
	}
	return int32(bp)
}

func (s *Server) Solve2015day14part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-14.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: getDistance(trimmed, 2503)}, nil
}

func (s *Server) Solve2015day14part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-14.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: runPoints(trimmed, 2503)}, nil
}
