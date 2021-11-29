package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func isTriangle(s1, s2, s3 int) bool {
	if s1+s2 <= s3 ||
		s2+s3 <= s1 ||
		s1+s3 <= s2 {
		return false
	}
	return true
}

func validTriangles(data string) int32 {
	count := int32(0)

	for _, elem := range strings.Split(data, "\n") {
		bits := strings.Fields(elem)
		s1, _ := strconv.Atoi(bits[0])
		s2, _ := strconv.Atoi(bits[1])
		s3, _ := strconv.Atoi(bits[2])

		if isTriangle(s1, s2, s3) {
			count++
		}
	}

	return count
}

func (s *Server) Solve2016day3part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-3.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(validTriangles(trimmed))}, nil
}

func (s *Server) Solve2016day3part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-3.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	total := int32(0)
	c1 := ""
	c2 := ""
	c3 := ""
	count := 0
	for _, line := range strings.Split(trimmed, "\n") {
		elems := strings.Fields(line)
		if count == 0 {
			c1 = elems[0]
			c2 = elems[1]
			c3 = elems[2]
		} else {
			c1 += " " + elems[0]
			c2 += " " + elems[1]
			c3 += " " + elems[2]
		}
		count++
		if count == 2 {
			total += validTriangles(c1 + "\n" + c2 + "\n" + c3)
			count = 0
		}
	}

	return &pb.SolveResponse{Answer: total}, nil
}
