package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getChecksum(str string) int32 {
	checksum := 0

	for _, line := range strings.Split(str, "\n") {
		minv := math.MaxInt32
		maxv := 0
		for _, digit := range strings.Fields(line) {
			num, err := strconv.Atoi(digit)
			if err != nil {
				log.Fatalf("Bad number: %v", digit)
			}

			if num < minv {
				minv = num
			}
			if num > maxv {
				maxv = num
			}
		}

		checksum += maxv - minv
	}

	return int32(checksum)
}

func (s *Server) Solve2017day2part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: getChecksum(trimmed)}, nil
}
