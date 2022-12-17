package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildMMap(data string) (map[string][]string, map[string]int) {
	mmap := make(map[string][]string)
	vals := make(map[string]int)

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			elems := strings.Fields(line)
			name := elems[1]
			bits := strings.Split(strings.Replace(elems[4], ";", "", -1), "=")
			val := getInt32(bits[1])

			vals[name] = val

			var follow []string
			for _, next := range elems[9:] {
				val := strings.Replace(next, ",", "", -1)
				follow = append(follow, val)
			}
			mmap[name] = follow
		}
	}

	return mmap, vals
}

type gasNode struct {
	cvalve    string
	active    map[string]bool
	remaining int
	sofar     int
	path      string
}

func (g gasNode) rep() string {
	return g.cvalve + fmt.Sprintf("%v-%v", g.active, g.remaining)
}

func computeGas(active map[string]bool, vals map[string]int) int {
	val := 0
	for key, act := range active {
		if act {
			val += vals[key]
		}
	}
	return val
}

func copyActive(active map[string]bool, add string) map[string]bool {
	nmap := make(map[string]bool)

	for key, val := range active {
		nmap[key] = val
	}
	if add != "" {
		nmap[add] = true
	}
	return nmap
}

func findDists(mmap map[string][]string, cv string) map[string]int {
	eMap := make(map[string]int)
	eMap[cv] = 0

	queue := []string{cv}

	count := 0
	for len(queue) > 0 {
		count++

		var nqueue []string

		for _, entry := range queue {
			nexts := mmap[entry]
			for _, n := range nexts {
				if _, ok := eMap[n]; !ok {
					nqueue = append(nqueue, n)
					eMap[n] = count
				}
			}
		}

		queue = nqueue
	}

	return eMap
}

func releaseGas(data string, minutes int) int {
	mmap, vals := buildMMap(data)

	currValve := "AA"
	bestValue := 0

	queue := []*gasNode{{cvalve: currValve, active: make(map[string]bool), remaining: 30, sofar: 0, path: "AA"}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		dists := findDists(mmap, head.cvalve)
		for node, dist := range dists {
			if _, ok := head.active[node]; !ok && (dist+1) < head.remaining {
				addition := vals[node] * ((head.remaining - dist) - 1)
				if addition > 0 {
					if head.sofar+addition > bestValue {
						bestValue = head.sofar + addition
					}
					queue = append(queue,
						&gasNode{
							cvalve:    node,
							active:    copyActive(head.active, node),
							remaining: head.remaining - (dist + 1),
							sofar:     head.sofar + addition,
							path:      head.path + fmt.Sprintf("-%v", node),
						})
				}
			}
		}

	}

	return bestValue
}

func releaseGasPair(data string, minutes int) int {
	mmap, vals := buildMMap(data)

	currValve := "AA"
	bestValue := 0

	queue := []*gasNode{{cvalve: currValve, active: make(map[string]bool), remaining: 30, sofar: 0, path: "AA"}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		dists := findDists(mmap, head.cvalve)
		for node, dist := range dists {
			if _, ok := head.active[node]; !ok && (dist+1) < head.remaining {
				addition := vals[node] * ((head.remaining - dist) - 1)
				if addition > 0 {
					if head.sofar+addition > bestValue {
						bestValue = head.sofar + addition
					}
					queue = append(queue,
						&gasNode{
							cvalve:    node,
							active:    copyActive(head.active, node),
							remaining: head.remaining - (dist + 1),
							sofar:     head.sofar + addition,
							path:      head.path + fmt.Sprintf("-%v", node),
						})
				}
			}
		}

	}

	return bestValue
}

func (s *Server) Solve2022day16part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-16.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(releaseGas(data, 30))}, nil
}
