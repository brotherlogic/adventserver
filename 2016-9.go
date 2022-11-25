package main

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func expandString(str string) int {
	nstr := ""

	pointer := 0
	inmarker := false
	marker := ""
	for pointer < len(str) {
		if str[pointer] == '(' {
			inmarker = true
			marker = ""
			pointer++
		} else if str[pointer] == ')' {
			inmarker = false
			elems := strings.Split(marker, "x")
			chars, _ := strconv.ParseInt(elems[0], 10, 32)
			reps, _ := strconv.ParseInt(elems[1], 10, 32)

			for i := 0; i < int(reps); i++ {
				for j := pointer + 1; j < pointer+1+int(chars); j++ {
					nstr += string(str[j])
				}
			}
			pointer += 1 + int(chars)
		} else if inmarker {
			marker += string(str[pointer])
			pointer++
		} else {
			nstr += string(str[pointer])
			pointer++
		}
	}
	log.Printf("%v -> %v", len(nstr), nstr)
	return len(nstr)
}

func searchString(str string) int64 {
	findex := strings.Index(str, "(")
	f2index := strings.Index(str, ")")
	if findex == -1 {
		return int64(len(str))
	}

	elems := strings.Split(str[findex+1:f2index], "x")
	chars, _ := strconv.ParseInt(elems[0], 10, 32)
	reps, _ := strconv.ParseInt(elems[1], 10, 32)

	return int64(findex) + reps*searchString(str[f2index+1:f2index+1+int(chars)]) + searchString(str[f2index+1+int(chars):])
}

func (s *Server) Solve2016day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-9.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(expandString(strings.TrimSpace(data)))}, nil
}

func (s *Server) Solve2016day9part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-9.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{BigAnswer: (searchString(strings.TrimSpace(data)))}, nil
}
