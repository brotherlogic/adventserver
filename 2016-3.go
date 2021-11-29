package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func isTriangle(s1, s2, s3 int) bool {
	if s1+s2 < s3 ||
		s2+s3 < s1 ||
		s1+s3 < s2 {
		return false
	}
	return true
}

func (s *Server) validTriangles(data string) int32 {
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

	return &pb.SolveResponse{Answer: int32(s.validTriangles(trimmed))}, nil
}
