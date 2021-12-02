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

func addRune(c rune, num int) string {
	val := (int(byte(c)-byte('a'))+num)%int(byte('z')-byte('a')+1) + int(byte('a'))
	return string(byte(val))
}

func trans(str string) string {
	bits := strings.Split(str, "[")
	parts := strings.Split(bits[0], "-")
	num, _ := strconv.Atoi(parts[len(parts)-1])

	fstr := ""
	for _, bit := range parts[:len(parts)-1] {
		for _, c := range bit {
			fstr += addRune(c, num)
		}
		fstr += " "
	}

	return strings.TrimSpace(fstr)
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

	ans := ""
	for _, str := range strings.Split(trimmed, "\n") {
		if strings.Contains(trans(str), "stor") {
			ans = trans(str)
		}
	}

	return &pb.SolveResponse{StringAnswer: ans}, nil
}
