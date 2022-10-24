package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	fullaunt = promauto.NewCounter(prometheus.CounterOpts{
		Name: "adventserver_fullaunt",
		Help: "The number of server requests",
	})
)

type properties struct {
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func findAunt(details string, known properties, fuzzy bool) int {
	lines := strings.Split(details, "\n")
	count := 0

	for i, line := range lines {
		nline := line[8:]

		if i < 10 {
			nline = line[7:]
		}

		found := true

		elems := strings.Fields(nline)
		for i := 0; i < len(elems); i += 2 {
			cat := elems[i][:len(elems[i])-1]
			values := elems[i+1]
			values = strings.TrimSuffix(values, ",")

			value, _ := strconv.ParseInt(values, 10, 32)

			switch cat {
			case "children":
				if value != int64(known.children) {
					found = false
				}
			case "cats":
				if (!fuzzy && value != int64(known.cats)) ||
					(fuzzy && value > int64(known.cats)) {
					found = false
				}
			case "samoyeds":
				if value != int64(known.samoyeds) {
					found = false
				}
			case "pomeranians":
				if (!fuzzy && value != int64(known.pomeranians)) ||
					(fuzzy && value < int64(known.pomeranians)) {
					found = false
				}
			case "akitas":
				if value != int64(known.akitas) {
					found = false
				}
			case "vizslas":
				if value != int64(known.vizslas) {
					found = false
				}
			case "goldfish":
				if (!fuzzy && value != int64(known.goldfish)) ||
					(fuzzy && value < int64(known.goldfish)) {
					found = false
				}
			case "trees":
				if (!fuzzy && value != int64(known.trees)) ||
					(fuzzy && value > int64(known.trees)) {
					found = false
				}
			case "cars":
				if value != int64(known.cars) {
					found = false
				}
			case "perfumes":
				if value != int64(known.perfumes) {
					found = false
				}
			}
		}

		if found {
			if !fuzzy {
				return i + 1
			}
			fullaunt.Inc()
			count++
		}
	}

	return count
}

func (s *Server) Solve2015day16part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-16.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	known := properties{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	return &pb.SolveResponse{Answer: int32(findAunt(trimmed, known, false))}, nil
}

func (s *Server) Solve2015day16part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-16.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	known := properties{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	return &pb.SolveResponse{Answer: int32(findAunt(trimmed, known, true))}, nil
}
