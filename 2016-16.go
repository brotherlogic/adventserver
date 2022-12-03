package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func flip(r rune) rune {
	if r == '0' {
		return '1'
	}
	return '0'
}

func reverseAndSwitch(s string) string {
	var runes []rune
	for _, r := range s {
		runes = append(runes, flip(r))
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = (runes[j]), (runes[i])
	}
	return string(runes)
}

func dragonExpand(in string) string {

	in2 := reverseAndSwitch(in)

	return in + "0" + in2
}

func dragonChecksum(in string) string {
	nstr := ""
	for i := 0; i < len(in); i += 2 {
		if in[i] == in[i+1] {
			nstr += "1"
		} else {
			nstr += "0"
		}
	}
	return nstr
}

func dragonRun(in string, size int) string {
	data := in
	for len(data) < size {
		data = dragonExpand(data)
	}

	rdata := data[:size]

	checksum := dragonChecksum(rdata)
	for len(checksum)%2 == 0 {
		checksum = dragonChecksum(checksum)
	}

	return checksum
}

func (s *Server) Solve2016day16part1(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{StringAnswer: (dragonRun("11100010111110100", 272))}, nil
}

func (s *Server) Solve2016day16part2(ctx context.Context) (*pb.SolveResponse, error) {

	return &pb.SolveResponse{StringAnswer: (dragonRun("11100010111110100", 35651584))}, nil
}
