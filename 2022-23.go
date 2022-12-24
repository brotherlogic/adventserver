package main

import (
	"fmt"
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type elfLocation struct {
	x, y int
}

type proposal struct {
	elfIndex   int
	newx, newy int
}

func findElves(data string) []*elfLocation {
	var ret []*elfLocation

	for y, line := range strings.Split(strings.TrimSpace(data), "\n") {
		for x, char := range line {
			if char == '#' {
				ret = append(ret, &elfLocation{x: x + 1, y: y + 1})
			}
		}
	}

	return ret
}

func buildPropos(elves []*elfLocation, start int) []*proposal {
	var props []*proposal

	elfLoc := make(map[string]bool)
	for _, elf := range elves {
		elfLoc[elf.getKey()] = true
	}

	for el, elf := range elves {
		for i := 0; i < 4; i++ {
			found := false

			count := 0
			if elfLoc[fmt.Sprintf("%v-%v", elf.x, elf.y-1)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x, elf.y+1)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y-1)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y+1)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y-1)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y)] {
				count++
			}
			if elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y+1)] {
				count++
			}

			if count != 0 {
				switch (i + start) % 4 {
				case 0:
					if !elfLoc[fmt.Sprintf("%v-%v", elf.x, elf.y-1)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y-1)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y-1)] {
						props = append(props, &proposal{newx: elf.x, newy: elf.y - 1, elfIndex: el})
						found = true
					}
				case 1:
					if !elfLoc[fmt.Sprintf("%v-%v", elf.x, elf.y+1)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y+1)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y+1)] {
						props = append(props, &proposal{newx: elf.x, newy: elf.y + 1, elfIndex: el})
						found = true
					}
				case 2:
					if !elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y+1)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x-1, elf.y-1)] {
						props = append(props, &proposal{newx: elf.x - 1, newy: elf.y, elfIndex: el})
						found = true
					}
				case 3:
					if !elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y+1)] &&
						!elfLoc[fmt.Sprintf("%v-%v", elf.x+1, elf.y-1)] {
						props = append(props, &proposal{newx: elf.x + 1, newy: elf.y, elfIndex: el})
						found = true
					}
				}
				if found {
					break
				}
			}
		}
	}

	return props
}

func (e *elfLocation) getKey() string {
	return fmt.Sprintf("%v-%v", e.x, e.y)
}

func (p *proposal) getKey() string {
	return fmt.Sprintf("%v-%v", p.newx, p.newy)
}

func printElves(elves []*elfLocation) string {
	minex, miney := math.MaxInt, math.MaxInt
	maxex, maxey := 0, 0

	for _, elf := range elves {
		minex = min(minex, elf.x)
		maxex = max(maxex, elf.x)

		miney = min(miney, elf.y)
		maxey = max(maxey, elf.y)
	}

	ret := ""
	for y := miney; y <= maxey; y++ {
		for x := minex; x <= maxex; x++ {
			found := false
			for _, elf := range elves {

				if elf.x == x && elf.y == y {
					ret += "#"
					found = true
				}
			}

			if !found {
				ret += "."
			}
		}
		ret += "\n"
	}

	return ret
}

func cleanProps(props []*proposal) []*proposal {
	propMap := make(map[string]int)

	for _, prop := range props {
		propMap[prop.getKey()]++
	}

	var nprops []*proposal
	for _, prop := range props {
		if propMap[prop.getKey()] == 1 {
			nprops = append(nprops, prop)
		}
	}

	return nprops
}

func elfSize(elves []*elfLocation) int {
	minex, miney := math.MaxInt, math.MaxInt
	maxex, maxey := 0, 0

	for _, elf := range elves {
		minex = min(minex, elf.x)
		maxex = max(maxex, elf.x)

		miney = min(miney, elf.y)
		maxey = max(maxey, elf.y)
	}

	return (maxex - minex + 1) * (maxey - miney + 1)
}

func runElves(data string, rounds int) int {
	elves := findElves(data)

	for i := 0; i < rounds; i++ {
		proposals := cleanProps(buildPropos(elves, i%4))
		if len(proposals) == 0 {
			return i + 1
		}
		for _, prop := range proposals {
			elves[prop.elfIndex].x = prop.newx
			elves[prop.elfIndex].y = prop.newy
		}
	}

	return elfSize(elves) - len(elves)
}

func (s *Server) Solve2022day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-23.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runElves(data, 10))}, nil
}

func (s *Server) Solve2022day23part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-23.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runElves(data, math.MaxInt))}, nil
}
