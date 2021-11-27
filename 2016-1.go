package main

import (
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

func runCoords(str string) int {
	coords := strings.Split(str, ",")
	heading := 0
	cx := 0
	cy := 0

	for _, cd := range coords {
		c := strings.TrimSpace(cd)
		char := c[0]
		switch char {
		case 'L':
			heading++
		case 'R':
			heading--
		}

		count, err := strconv.Atoi(c[1:len(c)])
		if err != nil {
			log.Fatalf("Bad convert: %v (%v)", err, c)
		}

		heading = heading % 4
		switch heading {
		case 0:
			cy += count
		case 1, -1:
			cx += count
		case 2, -2:
			cy -= count
		case 3, -3:
			cx -= count
		default:
			log.Fatalf("Cannot procecss heading %v", heading)
		}
	}

	return abs(cx) + abs(cy)
}

func (s *Server) solve2016day1part1(ctx context.Context) (int32, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-1.txt")
	if err != nil {
		return -1, err
	}

	return int32(runCoords(data)), nil
}
