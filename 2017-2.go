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

func evenDivide(str string) int32 {
	checksum := 0

	for _, line := range strings.Split(str, "\n") {
		for i, digit := range strings.Fields(line) {
			for j := i + 1; j < len(strings.Fields(line)); j++ {
				onum, err2 := strconv.Atoi(strings.Fields(line)[j])
				num, err := strconv.Atoi(digit)
				if err != nil || err2 != nil {
					log.Fatalf("Bad number: %v or %v", digit, strings.Fields(line)[j])
				}

				if num%onum == 0 {
					checksum += num / onum
				}
				if onum%num == 0 {
					checksum += onum / num
				}
			}
		}
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

func (s *Server) Solve2017day2part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: evenDivide(trimmed)}, nil
}
