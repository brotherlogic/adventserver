package main

import (
	"fmt"
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

func getIps(ctx context.Context, data string, m int64, log func(context.Context, string)) int64 {
	var ranges [][]int64

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			elems := strings.Split(line, "-")
			st, _ := strconv.ParseInt(elems[0], 10, 32)
			en, _ := strconv.ParseInt(elems[1], 10, 32)

			found := false
			for i := 0; i < len(ranges); i++ {
				// See if the start overlaps
				if (st) >= ranges[i][0] && (st) <= ranges[i][1] {
					if (en) > ranges[i][1] {
						ranges[i][1] = (en)
						found = true
					}
				}

				//End overlap
				if (en) >= ranges[i][0] && (en) <= ranges[i][1] {
					if (st) < ranges[i][0] {
						ranges[i][0] = (st)
						found = true
					}
				}

				//Envelops
				if (st) <= ranges[i][0] && (en) >= ranges[i][1] {
					ranges[i] = []int64{(st), (en)}
					found = true
				}

				if (st) >= ranges[i][0] && (en) <= ranges[i][1] {
					found = true
				}
			}
			if !found {
				ranges = append(ranges, []int64{(st), (en)})
			}

		}
	}

	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	log(ctx, fmt.Sprintf("%v", ranges))

	sumv := int64(0)
	for i := 0; i < len(ranges)-1; i++ {
		sumv += ranges[i+1][0] - ranges[i][1] - 1
	}

	sumv += m - ranges[len(ranges)-1][1]

	return sumv
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

	return &pb.SolveResponse{BigAnswer: (getIps(ctx, data, 4294967295, s.CtxLog))}, nil
}
