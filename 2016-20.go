package main

import (
	"fmt"
	"math/big"
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

func runMerge(p1, p2 []int64) ([]int64, bool) {
	if p1[0] <= p2[1] && p1[0] >= p2[0] {
		if p1[1] >= p2[1] {
			return []int64{p2[0], p1[1]}, true
		}
		return []int64{p2[0], p2[1]}, true
	}

	if p1[1] <= p2[1] && p1[1] >= p2[0] {
		if p1[0] <= p2[0] {
			return []int64{p1[0], p2[1]}, true
		}
		return []int64{p2[0], p2[1]}, true
	}

	if p1[0] <= p2[0] && p1[1] >= p2[1] {
		return []int64{p1[0], p1[1]}, true
	}

	if p2[0] <= p1[0] && p2[1] >= p1[1] {
		return []int64{p2[0], p2[1]}, true
	}

	return []int64{}, false
}

func getIps(ctx context.Context, data string, m int64, rlog func(context.Context, string)) string {
	var ranges [][]int64

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			elems := strings.Split(line, "-")
			st, _ := strconv.ParseInt(elems[0], 10, 64)
			en, _ := strconv.ParseInt(elems[1], 10, 64)

			ranges = append(ranges, []int64{st, en})
		}
	}

	merge := true
	for merge {
		merge = false
		for p1 := 0; p1 < len(ranges); p1++ {
			for p2 := p1 + 1; p2 < len(ranges); p2++ {
				if val, ok := runMerge(ranges[p1], ranges[p2]); ok {
					ranges[p1] = val
					ranges = append(ranges[:p2], ranges[p2+1:]...)
					merge = true
				}
			}
		}
	}

	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	rlog(ctx, fmt.Sprintf("%v", ranges))

	sumv := big.NewInt(0)
	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i+1][0] != ranges[i][1] {
			if ranges[i+1][0] > ranges[i][1] {
				sumv.Add(sumv, big.NewInt(ranges[i+1][0]-ranges[i][1]-int64(1)))
			}
		}
	}

	sumv.Add(sumv, big.NewInt(m-ranges[len(ranges)-1][1]))

	return fmt.Sprintf("%v", sumv)
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
				// See if the start overlaps
				if int(st) >= ranges[i][0] && int(st) <= ranges[i][1] {
					if int(en) > ranges[i][1] {
						ranges[i][1] = int(en)
						found = true
					}
				}

				//End overlap
				if int(en) >= ranges[i][0] && int(en) <= ranges[i][1] {
					if int(st) < ranges[i][0] {
						ranges[i][0] = int(st)
						found = true
					}
				}

				//Envelops
				if int(st) <= ranges[i][0] && int(en) >= ranges[i][1] {
					ranges[i] = []int{int(st), int(en)}
					found = true
				}

				if int(st) >= ranges[i][0] && int(en) <= ranges[i][1] {
					found = true
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
		if ranges[i+1][0] > ranges[i][1]+1 {
			return ranges[i][1] + 1
		}
	}

	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][1] < ranges[j][1]
	})

	return ranges[len(ranges)-1][1] + 1
}

func (s *Server) Solve2016day20part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-20.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getLowIp(data))}, nil
}

func (s *Server) Solve2016day20part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-20.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: (getIps(ctx, data, 4294967295, s.CtxLog))}, nil
}
