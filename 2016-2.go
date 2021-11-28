package main

import (
	"log"
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func procCode(code string, sx, sy int) int32 {
	res := 0
	keypad := [3][3]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	elems := strings.Split(code, "\n")
	multiplier := int(math.Pow10(len(elems) - 1))
	for _, elem := range elems {
		for _, c := range elem {
			switch c {
			case 'U':
				sy++
			case 'D':
				sy--
			case 'L':
				sx--
			case 'R':
				sx++
			default:
				log.Fatalf("Cannot process %v", c)
			}

			if sx < 0 {
				sx = 0
			}
			if sx > 2 {
				sx = 2
			}
			if sy < 0 {
				sy = 0
			}
			if sy > 2 {
				sy = 2
			}
		}

		res += multiplier * keypad[sx][sy]
		multiplier /= 10
	}
	return int32(res)
}

func procCompCode(code string, sx, sy int) string {
	res := ""
	keypad := [][]string{}
	keypad = append(keypad, []string{"Z", "Z", "Z", "Z", "Z", "Z", "Z"})
	keypad = append(keypad, []string{"Z", "Z", "Z", "5", "Z", "Z", "Z"})
	keypad = append(keypad, []string{"Z", "Z", "A", "6", "2", "Z", "Z"})
	keypad = append(keypad, []string{"Z", "D", "B", "7", "3", "1", "Z"})
	keypad = append(keypad, []string{"Z", "Z", "C", "8", "4", "Z", "Z"})
	keypad = append(keypad, []string{"Z", "Z", "Z", "9", "Z", "Z", "Z"})
	keypad = append(keypad, []string{"Z", "Z", "Z", "Z", "Z", "Z", "Z"})

	elems := strings.Split(code, "\n")
	for _, elem := range elems {
		for _, c := range elem {
			switch c {
			case 'U':
				if keypad[sx][sy+1] != "Z" {
					sy++
				}
			case 'D':
				if keypad[sx][sy-1] != "Z" {
					sy--
				}
			case 'L':
				if keypad[sx-1][sy] != "Z" {
					sx--
				}
			case 'R':
				if keypad[sx+1][sy] != "Z" {
					sx++
				}
			default:
				log.Fatalf("Cannot process %v", c)
			}

		}

		res += keypad[sx][sy]
	}
	return res
}

func (s *Server) Solve2016day2part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: procCode(trimmed, 1, 1)}, nil
}

func (s *Server) Solve2016day2part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-2.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{StringAnswer: procCompCode(trimmed, 1, 3)}, nil
}
