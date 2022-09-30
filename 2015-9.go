package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func computeBestDistance(details string) int64 {
	places, distance := buildDistanceMap(details)

	best := runCompute(make([]string, 0), places, distance)

	return best
}

func runCompute(sofar, places []string, distance map[string]int64) int64 {
	if len(places) == 0 {
		dist := int64(0)
		for i := 0; i < len(sofar)-1; i++ {
			dist += distance[fmt.Sprintf("%v|%v", sofar[i], sofar[i+1])]
		}
		return dist
	}

	best := int64(math.MaxInt64)
	for _, place := range places {
		var nsofar []string
		copy(nsofar, sofar)
		sofar = append(sofar, place)
		var nplace []string
		for _, p := range places {
			if p != place {
				nplace = append(nplace, p)
			}
		}

		nbest := runCompute(sofar, nplace, distance)
		if nbest < best {
			best = nbest
		}
	}

	return best
}

func buildDistanceMap(details string) ([]string, map[string]int64) {
	retMap := make(map[string]int64)
	var places []string
	for _, line := range strings.Split(details, "\n") {
		elems := strings.Fields(line)
		dist, _ := strconv.ParseInt(elems[4], 10, 32)
		retMap[fmt.Sprintf("%v|%v", elems[0], elems[2])] = dist
		retMap[fmt.Sprintf("%v|%v", elems[2], elems[0])] = dist

		found0 := false
		found2 := false
		for _, place := range places {
			found0 = found0 || place == elems[0]
			found2 = found2 || place == elems[2]
		}
		if !found0 {
			places = append(places, elems[0])
		}
		if !found2 {
			places = append(places, elems[2])
		}
	}

	return places, retMap
}

func (s *Server) Solve2015day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-9.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{BigAnswer: computeBestDistance(trimmed)}, nil
}
