package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findOptimal(s string, me bool) int {

	mapping := make(map[string]int)
	var names []string

	for _, line := range strings.Split(s, "\n") {
		elems := strings.Fields(line)
		n1 := elems[0]
		n2 := strings.ReplaceAll(elems[10], ".", "")

		happiness, _ := strconv.ParseInt(elems[3], 10, 32)
		if elems[2] == "lose" {
			happiness = 0 - happiness
		}

		mapping[n1+"|"+n2] = int(happiness)

		found1, found2 := false, false
		for _, n := range names {
			found1 = found1 || n == n1
			found2 = found2 || n == n2
		}

		if !found1 {
			names = append(names, n1)
		}
		if !found2 {
			names = append(names, n2)
		}
	}

	if me {
		for _, n := range names {
			mapping[fmt.Sprintf("%v|%v", "ME", n)] = 0
			mapping[fmt.Sprintf("%v|%v", n, "ME")] = 0
		}
		names = append(names, "ME")
	}

	return findSeating(mapping, names)
}

func findSeating(mapp map[string]int, names []string) int {
	b := 0
	for _, name := range names {
		var nname []string
		for _, n := range names {
			if n != name {
				nname = append(nname, n)
			}
		}
		bn := findSeats([]string{name}, nname, mapp)
		if bn > b {
			b = bn
		}
	}
	return b
}

func findSeats(pot []string, names []string, mapp map[string]int) int {
	if len(names) == 0 {

		happiness := 0
		for i := 0; i < len(pot); i++ {
			if i == 0 {
				happiness += mapp[fmt.Sprintf("%v|%v", pot[i], pot[len(pot)-1])]
			} else {
				happiness += mapp[fmt.Sprintf("%v|%v", pot[i], pot[i-1])]

			}

			if i == len(pot)-1 {
				happiness += mapp[fmt.Sprintf("%v|%v", pot[i], pot[0])]
			} else {
				happiness += mapp[fmt.Sprintf("%v|%v", pot[i], pot[i+1])]
			}
		}

		return happiness
	}

	b := 0
	for _, name := range names {
		var nname []string
		for _, n := range names {
			if n != name {
				nname = append(nname, n)
			}
		}
		npot := pot
		npot = append(npot, name)
		bn := findSeats(npot, nname, mapp)
		if bn > b {
			b = bn
		}
	}
	return b
}

func (s *Server) Solve2015day13part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-13.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(findOptimal(trimmed, false))}, nil
}
func (s *Server) Solve2015day13part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-13.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(findOptimal(trimmed, true))}, nil
}
