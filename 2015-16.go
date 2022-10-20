package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
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

func findAunt(details string, known properties) int {
	lines := strings.Split(details, "\n")

	for i, line := range lines {
		nline := line[8:]

		if i < 10 {
			nline = line[7:]
		}

		found := false

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
				if value != int64(known.cats) {
					found = false
				}
			case "samoyeds":
				if value != int64(known.samoyeds) {
					found = false
				}
			case "pomeranians":
				if value != int64(known.pomeranians) {
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
				if value != int64(known.goldfish) {
					found = false
				}
			case "trees":
				if value != int64(known.trees) {
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
			return i + 1
		}
	}

	return 0
}

func (s *Server) Solve2015day16part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-16.txt")
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

	return &pb.SolveResponse{Answer: int32(findAunt(trimmed, known))}, nil
}
