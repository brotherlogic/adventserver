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
	gqlen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2022_16_qlen",
		Help: "The number of server requests",
	})
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
	cvalve     string
	bvalve     string
	active     map[string]bool
	remaining  int
	bremaining int
	sofar      int
	path       string
	bpath      string
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

func releaseGas(data string, minutes int) (int, map[string]int) {
	mmap, vals := buildMMap(data)
	fmap := make(map[string]int)

	currValve := "AA"
	bestValue := 0

	queue := []*gasNode{{cvalve: currValve, active: make(map[string]bool), remaining: minutes, sofar: 0, path: "AA"}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		dists := findDists(mmap, head.cvalve)
		for node, dist := range dists {
			if _, ok := head.active[node]; !ok && (dist+1) < head.remaining {
				addition := vals[node] * ((head.remaining - dist) - 1)
				if addition > 0 {
					fmap[fmt.Sprintf("%v-%v", head.path, node)] = head.sofar + addition
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

	return bestValue, fmap
}

func gasOverlap(ar, br map[string]bool) bool {
	if len(ar) > len(br) {
		for a := range ar {
			if br[a] {
				return true
			}
		}
	} else {
		for b := range br {
			if ar[b] {
				return true
			}
		}
	}

	return false
}

func releaseGasSimple(data string, minutes int) int {
	_, fmap := releaseGas(data, minutes)
	var keys []string
	kmap := make(map[string]map[string]bool)

	for key := range fmap {
		keys = append(keys, key)

		nmap := make(map[string]bool)
		for _, val := range strings.Split(key, "-") {
			if val != "AA" {
				nmap[val] = true
			}
		}
		kmap[key] = nmap
	}

	best := 0
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if !gasOverlap(kmap[keys[i]], kmap[keys[j]]) && fmap[keys[i]]+fmap[keys[j]] > best {
				best = fmap[keys[i]] + fmap[keys[j]]
			}
		}
	}

	return best
}

func releaseGasPair(data string, minutes int) int {
	mmap, vals := buildMMap(data)

	currValveMe := "AA"
	currValveBear := "AA"
	bestValue := 0

	queue := []*gasNode{{cvalve: currValveMe, bvalve: currValveBear, active: make(map[string]bool), remaining: minutes, bremaining: minutes, sofar: 0, path: "AA", bpath: "AA"}}

	seen := 0
	for len(queue) > 0 {
		gqlen.Set(float64(len(queue)))
		head := queue[0]
		queue = queue[1:]

		dists := findDists(mmap, head.cvalve)
		bdists := findDists(mmap, head.bvalve)
		for node, dist := range dists {
			if _, ok := head.active[node]; !ok && (dist+1) < head.remaining {
				addition := vals[node] * ((head.remaining - dist) - 1)
				baddition := vals[node] * ((head.bremaining - bdists[node]) - 1)
				if addition > 0 {

					seen++
					if head.remaining >= head.bremaining {
						if head.sofar+addition > bestValue {
							bestValue = head.sofar + addition
						}

						queue = append(queue,
							&gasNode{
								cvalve:     node,
								bvalve:     head.bvalve,
								active:     copyActive(head.active, node),
								remaining:  head.remaining - (dist + 1),
								bremaining: head.bremaining,
								sofar:      head.sofar + addition,
								path:       head.path + fmt.Sprintf("-%v", node),
								bpath:      head.bpath,
							})
					}
					if head.bremaining >= head.remaining && len(head.bpath) < len(head.path) {
						if head.sofar+baddition > bestValue {
							bestValue = head.sofar + baddition
						}
						queue = append(queue,
							&gasNode{
								bvalve:     node,
								cvalve:     head.cvalve,
								active:     copyActive(head.active, node),
								bremaining: head.bremaining - (bdists[node] + 1),
								remaining:  head.remaining,
								sofar:      head.sofar + baddition,
								bpath:      head.bpath + fmt.Sprintf("-%v", node),
								path:       head.path,
							})
					}
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

	ans, _ := releaseGas(data, 30)
	return &pb.SolveResponse{Answer: int32(ans)}, nil
}

func (s *Server) Solve2022day16part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-16.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(releaseGasSimple(data, 26))}, nil
}
