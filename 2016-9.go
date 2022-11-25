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
func (s *Server) Solve2016day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-9.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(expandString(strings.TrimSpace(data)))}, nil
}
