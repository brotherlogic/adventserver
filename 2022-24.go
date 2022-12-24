package main

import (
	"fmt"
	"log"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	blen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2022_24_qlen",
		Help: "The number of server requests",
	})
)

type blizzNode struct {
	cycle  int
	px, py int
}

type blizzard struct {
	cycle      int
	bx, by, bd []int
	bmx, bmy   int
}

func (b *blizzard) next() *blizzard {
	nblizzard := &blizzard{
		cycle: b.cycle + 1,
		bmx:   b.bmx,
		bmy:   b.bmy,
		bx:    make([]int, 0),
		by:    make([]int, 0),
		bd:    make([]int, 0),
	}

	for i := range b.bx {
		nblizzard.bd = append(nblizzard.bd, b.bd[i])
		switch b.bd[i] {
		case 0:
			nblizzard.by = append(nblizzard.by, b.by[i])
			if b.bx[i] == b.bmx-2 {
				nblizzard.bx = append(nblizzard.bx, 1)
			} else {
				nblizzard.bx = append(nblizzard.bx, b.bx[i]+1)
			}
		case 2:
			nblizzard.by = append(nblizzard.by, b.by[i])
			if b.bx[i] == 1 {
				nblizzard.bx = append(nblizzard.bx, b.bmx-2)
			} else {
				nblizzard.bx = append(nblizzard.bx, b.bx[i]-1)
			}
		case 1:
			nblizzard.bx = append(nblizzard.bx, b.bx[i])
			if b.by[i] == b.bmy-2 {
				nblizzard.by = append(nblizzard.by, 1)
			} else {
				nblizzard.by = append(nblizzard.by, b.by[i]+1)
			}
		case 3:
			nblizzard.bx = append(nblizzard.bx, b.bx[i])
			if b.by[i] == 1 {
				nblizzard.by = append(nblizzard.by, b.bmy-2)
			} else {
				nblizzard.by = append(nblizzard.by, b.by[i]-1)
			}
		}
	}

	return nblizzard
}

func buildBlizzard(data string) *blizzard {
	blizzard := &blizzard{cycle: 1, bx: make([]int, 0), by: make([]int, 0), bd: make([]int, 0)}

	sx, sy := getBlizzSize(data)
	blizzard.bmx = sx
	blizzard.bmy = sy

	for y, line := range strings.Split(strings.TrimSpace(data), "\n") {
		for x, char := range strings.TrimSpace(line) {
			if char != '#' && char != '.' {
				blizzard.bx = append(blizzard.bx, x)
				blizzard.by = append(blizzard.by, y)
				switch char {
				case '<':
					blizzard.bd = append(blizzard.bd, 2)
				case 'v':
					blizzard.bd = append(blizzard.bd, 1)
				case '>':
					blizzard.bd = append(blizzard.bd, 0)
				case '^':
					blizzard.bd = append(blizzard.bd, 3)
				default:
					log.Fatalf("WHAT WHAT WHAT: %v", string(char))
				}
			}
		}
	}

	return blizzard
}

func getBlizzSize(data string) (int, int) {
	elems := strings.Split(strings.TrimSpace(data), "\n")
	return len(elems[0]), len(elems)
}

func (b *blizzard) occupied(x, y int) bool {
	if x == b.bmx-2 && y == b.bmy-1 {
		return false
	}
	if x == 1 && y == 0 {
		return false
	}

	if x == 0 || y <= 0 || x == b.bmx-1 || y == b.bmy-1 {
		return true
	}

	for i := range b.bx {
		if b.bx[i] == x && b.by[i] == y {
			return true
		}
	}

	return false
}

func (b *blizzard) print(px, py int) string {
	ret := ""
	for y := 0; y < b.bmy; y++ {
		for x := 0; x < b.bmx; x++ {
			if x == 0 || y == 0 {
				if x == 1 {
					if x == px && y == py {
						ret += "E"
					} else {
						ret += "."
					}
				} else {
					ret += "#"
				}
			} else if x == b.bmx-1 || y == b.bmy-1 {
				if x == b.bmx-2 {
					if x == px && y == py {
						ret += "E"
					} else {
						ret += "."
					}
				} else {
					ret += "#"
				}
			} else {
				if x == px && y == py {
					ret += "E"
				} else {
					count := 0
					val := ""
					for i := range b.bx {
						if b.bx[i] == x && b.by[i] == y {
							count++
							switch b.bd[i] {
							case 0:
								val = ">"
							case 1:
								val = "v"
							case 2:
								val = "<"
							case 3:
								val = "^"
							}
						}
					}
					if count == 0 {
						ret += "."
					} else {
						if count == 1 {
							ret += val
						} else {
							ret += fmt.Sprintf("%v", count)
						}
					}
				}
			}
		}
		ret += "\n"
	}

	if b.occupied(px, py+1) {
		ret += "OCCUPY\n"
	} else {
		ret += "FREE: %v\n"
	}

	return ret
}

func (p *blizzNode) next(b *blizzard) []*blizzNode {
	var ret []*blizzNode

	// Wait
	if !b.occupied(p.px, p.py) {
		ret = append(ret, &blizzNode{px: p.px, py: p.py, cycle: p.cycle + 1})
	}

	if !b.occupied(p.px+1, p.py) {
		ret = append(ret, &blizzNode{px: p.px + 1, py: p.py, cycle: p.cycle + 1})
	}
	if !b.occupied(p.px-1, p.py) {
		ret = append(ret, &blizzNode{px: p.px - 1, py: p.py, cycle: p.cycle + 1})
	}
	if !b.occupied(p.px, p.py+1) {
		ret = append(ret, &blizzNode{px: p.px, py: p.py + 1, cycle: p.cycle + 1})
	}
	if !b.occupied(p.px, p.py-1) {
		ret = append(ret, &blizzNode{px: p.px, py: p.py - 1, cycle: p.cycle + 1})
	}

	return ret
}

func (b *blizzNode) rep() string {
	return fmt.Sprintf("%v-%v-%v", b.cycle, b.px, b.py)
}

func runBlizzardMaze(blizz *blizzard, end int) (int, *blizzard) {
	bMaze := make(map[int]*blizzard)

	bx, by := blizz.bmx, blizz.bmy
	bMaze[0] = blizz
	seen := make(map[string]bool)

	queue := []*blizzNode{{cycle: 0, px: 1, py: 0}}
	if end == 0 {
		queue[0].px, queue[0].py = bx-2, by-1
	}

	for len(queue) > 0 {
		blen.Set(float64(len(queue)))
		head := queue[0]

		queue = queue[1:]

		if end == 1 {
			if head.px == bx-2 && head.py == by-1 {
				return head.cycle, bMaze[head.cycle]
			}
		} else {
			if head.px == 1 && head.py == 0 {
				return head.cycle, bMaze[head.cycle]
			}
		}

		if _, ok := bMaze[head.cycle+1]; !ok {
			bMaze[head.cycle+1] = bMaze[head.cycle].next()
		}

		nexts := head.next(bMaze[head.cycle+1])
		for _, n := range nexts {
			if _, ok := seen[n.rep()]; !ok {
				queue = append(queue, n)
				seen[n.rep()] = true
			}
		}
	}

	return 0, nil
}

func (s *Server) Solve2022day24part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-24.txt")
	if err != nil {
		return nil, err
	}

	res, _ := runBlizzardMaze(buildBlizzard(data), 1)
	return &pb.SolveResponse{Answer: int32(res)}, nil
}

func (s *Server) Solve2022day24part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-24.txt")
	if err != nil {
		return nil, err
	}

	blizz := buildBlizzard(data)
	res1, blizz1 := runBlizzardMaze(blizz, 1)
	res2, blizz2 := runBlizzardMaze(blizz1, 0)
	res3, _ := runBlizzardMaze(blizz2, 1)
	return &pb.SolveResponse{Answer: int32(res1 + res2 + res3)}, nil
}
