package main

import (
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isRealRoom(str string) int {
	bits := strings.Split(str, "[")
	lcount := make(map[string]int)

	parts := strings.Split(bits[0], "-")
	for i, p := range parts {
		if i < len(parts)-1 {
			for _, r := range p {
				lcount[string(r)]++
			}
		}
	}

	var counts []int
	for _, count := range lcount {
		found := false
		for _, c := range counts {
			if c == count {
				found = true
			}
		}
		if !found {
			counts = append(counts, count)
		}
	}
	sort.Ints(counts)

	fstr := ""
	for _, c := range counts {
		nstr := ""
		for r, co := range lcount {
			if co == c {
				nstr += string(r)
			}
		}
		nstr = sortString(nstr)
		fstr = nstr + fstr
	}

	b2 := strings.Split(bits[1], "]")
	if strings.HasPrefix(fstr, b2[0]) {
		num, _ := strconv.Atoi(parts[len(parts)-1])
		return num
	}

	return 0
}

func (s *Server) Solve2017day4part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-4.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	count := 0
	for _, str := range strings.Split(trimmed, "\n") {
		count += isRealRoom(str)
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
