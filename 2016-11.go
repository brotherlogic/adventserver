package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	states = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day16_searches",
		Help: "The number of server requests",
	})
)

type state struct {
	floors   map[int][]string
	elevator int
	moves    int
	path     string
}

func (s state) getRep() string {
	//ret := fmt.Sprintf("E%v", s.elevator)
	ret := ""
	for floor, list := range s.floors {
		ret += fmt.Sprintf("%v:", floor)

		cmap := make(map[byte]int)
		for _, elem := range list {
			cmap[elem[0]]++
		}

		ms := 0
		gs := 0
		cs := 0
		for _, elem := range list {
			if elem[1] == 'G' {
				if val := cmap[elem[0]]; val == 2 {
					cs++
				} else {
					gs++
				}
			}

			if elem[1] == 'M' {
				if val := cmap[elem[0]]; val < 2 {
					ms++
				}
			}
		}

		ret += fmt.Sprintf("%v-%v-%v", ms, gs, cs)
	}

	return ret
}

func buildFloors(data string) state {
	floors := make(map[int][]string)

	for i, line := range strings.Split(data, "\n") {
		cline := strings.ReplaceAll(strings.ReplaceAll(line, ",", ""), ".", "")
		elems := strings.Fields(cline)
		for j, elem := range elems {
			if elem == "microchip" {
				parts := strings.Split(elems[j-1], "-")
				floors[i+1] = append(floors[i+1], fmt.Sprintf("%v%v", string(parts[0][0]), "M"))
			} else if elem == "generator" {
				floors[i+1] = append(floors[i+1], fmt.Sprintf("%v%v", string(elems[j-1][0]), "G"))
			}
		}
	}

	return state{floors: floors, elevator: 1}
}

func isWinner(st state) bool {
	for floor := 1; floor < 4; floor++ {
		if len(st.floors[floor]) > 0 {
			return false
		}
	}

	return true
}

func (s state) copy(pickups ...string) state {
	ns := state{elevator: s.elevator, moves: s.moves, floors: make(map[int][]string), path: s.path}
	for i := 1; i <= 4; i++ {
		var nf []string
		for _, elem := range s.floors[i] {
			found := false
			for _, p := range pickups {
				if p == elem {
					found = true
				}
			}

			if !found {
				nf = append(nf, elem)
			}
		}

		ns.floors[i] = nf
	}

	return ns
}

func isAllowed(base []string, adds ...string) bool {
	mapG := make(map[byte]bool)
	mapC := make(map[byte]bool)

	for _, str := range base {
		if strings.HasSuffix(str, "G") {
			mapG[str[0]] = true
		} else {
			mapC[str[0]] = true
		}
	}

	for _, str := range adds {
		if strings.HasSuffix(str, "G") {
			mapG[str[0]] = true
		} else {
			mapC[str[0]] = true
		}
	}

	if len(mapG) == 0 {
		return true
	}

	for str, _ := range mapC {
		if _, ok := mapG[str]; !ok {
			return false
		}
	}

	return true
}

func isLegalMove(nstate state, nfloor int, pickups ...string) (state, bool) {
	adjuster := 1
	if nfloor < nstate.elevator {
		adjuster = -1
	}
	for floor := nstate.elevator; floor != nfloor; floor += adjuster {
		if !isAllowed(nstate.floors[floor], pickups...) {
			return nstate, false
		}
	}

	if !isAllowed(nstate.floors[nfloor], pickups...) {
		return nstate, false
	}

	nstate.floors[nfloor] = append(nstate.floors[nfloor], pickups...)
	nstate.path += fmt.Sprintf("%v->%v;", strings.Join(pickups, ","), nfloor)
	nstate.moves += abs(nstate.elevator - nfloor)
	nstate.elevator = nfloor

	return nstate, true
}

func getBest(queue []state) (state, []state) {
	return queue[0], queue[1:]
}

func runFloorSearch(queue []state) (int, string) {

	seen := make(map[string]bool)
	var head state
	count := 0
	for len(queue) > 0 {
		states.Inc()
		count++
		head, queue = getBest(queue)
		if _, ok := seen[head.getRep()]; ok {
			continue
		}
		seen[head.getRep()] = true
		if isWinner(head) {
			return head.moves, head.path
		}
		for _, pickup1 := range head.floors[head.elevator] {
			// Run the single pickup first
			for nfloor := 1; nfloor <= 4; nfloor++ {
				if nfloor != head.elevator {
					nstate := head.copy(pickup1)
					_, found := seen[nstate.getRep()]
					if !found {
						nstate, ok := isLegalMove(nstate, nfloor, pickup1)
						if ok {
							queue = append(queue, nstate)
						}
					}
				}
			}
		}

		for _, pickup1 := range head.floors[head.elevator] {
			for _, pickup2 := range head.floors[head.elevator] {
				if pickup1 != pickup2 {
					// Run the single pickup first
					for nfloor := 1; nfloor <= 4; nfloor++ {
						if nfloor != head.elevator {
							nstate := head.copy(pickup1, pickup2)
							_, found := seen[nstate.getRep()]
							if !found {
								nstate, ok := isLegalMove(nstate, nfloor, pickup1, pickup2)
								if ok {
									queue = append(queue, nstate)
								}
							}
						}
					}
				}
			}
		}
	}

	return 0, ""
}

func findFloors(data string, add bool) (int, string, string) {
	floors := buildFloors(data)
	if add {
		floors.floors[1] = append(floors.floors[1], []string{"eG", "eM", "dG", "dM"}...)
	}

	str := fmt.Sprintf("%+v", floors)

	r, a := runFloorSearch([]state{floors})
	return r, a, str
}

func (s *Server) Solve2016day11part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-11.txt")
	if err != nil {
		return nil, err
	}

	res, path, start := findFloors(data, false)
	s.CtxLog(ctx, fmt.Sprintf("PATH %v -> %v", path, start))

	return &pb.SolveResponse{Answer: int32(res)}, nil
}

func (s *Server) Solve2016day11part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-11.txt")
	if err != nil {
		return nil, err
	}

	res, path, start := findFloors(data, true)
	s.CtxLog(ctx, fmt.Sprintf("PATH %v -> %v", path, start))

	return &pb.SolveResponse{Answer: int32(res)}, nil
}
