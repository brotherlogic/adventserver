package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func findNextPassword(s string) string {
	nstr := incrementString(s)
	for !isValidSantaPassword(nstr) {
		nstr = incrementString(nstr)
	}
	return nstr
}

func incrementString(s string) string {
	nstr := ""
	carry := true
	for i := len(s) - 1; i >= 0; i-- {
		if carry {
			if s[i] == 'z' {
				nstr = "a" + nstr
				carry = true
			} else {
				nstr = string(s[i]+1) + nstr
				carry = false
			}
		} else {
			nstr = string(s[i]) + nstr
		}
	}
	return nstr
}

func isValidSantaPassword(s string) bool {
	t := hasTrio(s)
	nvc := noBadChars(s)
	tp := twoPairs(s)

	return t && nvc && tp
}

func hasTrio(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			return true
		}
	}
	return false
}

func noBadChars(s string) bool {
	for _, c := range s {
		if c == 'l' || c == 'i' || c == 'o' {
			return false
		}
	}
	return true
}

func twoPairs(s string) bool {
	rcount := make(map[byte]int)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			rcount[s[i]]++
		}
	}
	return len(rcount) >= 2
}

func (s *Server) Solve2015day11part1(ctx context.Context) (*pb.SolveResponse, error) {
	start := "hepxcrrq"

	return &pb.SolveResponse{StringAnswer: findNextPassword(start)}, nil
}
