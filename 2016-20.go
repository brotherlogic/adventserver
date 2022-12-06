package main

import (
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func ipOverlap(comp []int, st, en int) ([]int, bool) {
	if st >= comp[0] && st <= comp[1] && en > comp[1] {
		return []int{comp[0], en}, true
	}

	if en <= comp[1] && en >= comp[0] && st < comp[0] {
		return []int{st, comp[1]}, true
	}

	return []int{}, false
}

func getLowIp(data string) int {
	var ranges [][]int

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			elems := strings.Split(line, "-")
			st, _ := strconv.ParseInt(elems[0], 10, 32)
			en, _ := strconv.ParseInt(elems[1], 10, 32)

			found := false
			for i := 0; i < len(ranges); i++ {
				if val, ok := ipOverlap(ranges[i], int(st), int(en)); ok {
					ranges[i] = val
					found = true
					break
				}
			}
			if !found {
				ranges = append(ranges, []int{int(st), int(en)})
			}
		}
	}

	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i][1] < ranges[i+1][0]-1 {
			return ranges[i][1] + 1
		}
	}

	return ranges[len(ranges)-1][1] + 1
}

func (s *Server) Solve2016day20part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-20.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getLowIp(data))}, nil
}
