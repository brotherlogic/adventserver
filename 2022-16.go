package main

import (
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

func releaseGas(data string, minutes int) int {
	mmap, vals := buildMMap(data)

	queue := []gasNode{{cvalve: "AA", active: make(map[string]bool), remaining: minutes, sofar: 0}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.remaining == 0 {
			return head.sofar
		}

		if vals[head.cvalve] > 0 && !head.active[head.cvalve] {
			queue = append(queue, gasNode{
				cvalve:    head.cvalve,
				active:    copyActive(head.active, head.cvalve),
				remaining: head.remaining - 1,
				sofar:     head.sofar + computeGas(head.active, vals)})
		}

		for _, next := range mmap[head.cvalve] {
			queue = append(queue, gasNode{
				cvalve:    next,
				active:    copyActive(head.active, ""),
				remaining: head.remaining - 1,
				sofar:     head.sofar + computeGas(head.active, vals),
			})
		}
	}

	return -1
}

func (s *Server) Solve2022day16part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-16.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(releaseGas(data, 30))}, nil
}
