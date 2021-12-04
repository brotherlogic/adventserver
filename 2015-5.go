package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func CountVowels(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func CountRow(str string) int {
	row_count := 1
	best := 0
	current := str[0]

	for i := 1; i < len(str); i++ {
		if str[i] == current {
			row_count++
			best = Max(best, row_count)
		} else {
			row_count = 1
		}
		current = str[i]
	}

	return best
}

func CountMaxNonOverlapping(str string) int {
	counts := make(map[string]int)
	for i := 0; i < len(str)-1; i++ {
		pair := string(str[i : i+2])
		safe := true

		if i < len(str)-2 {
			if str[i] == str[i+1] && str[i+1] == str[i+2] {
				i++
			}
		}

		if safe {
			counts[pair]++
		}
	}

	max_v := 0
	for _, value := range counts {
		max_v = Max(value, max_v)
	}
	return max_v
}

func RepeatWithMiddle(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func IsNiceAlso(str string) bool {
	return CountMaxNonOverlapping(str) >= 2 && RepeatWithMiddle(str)
}

func IsNice(str string) bool {
	bad_str := []string{"ab", "cd", "pq", "xy"}
	for i := 0; i < len(bad_str); i++ {
		if strings.Contains(str, bad_str[i]) {
			return false
		}
	}

	return CountVowels(str) >= 3 && CountRow(str) >= 2
}

func (s *Server) Solve2015day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-5.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	count := 0
	for _, str := range strings.Split(trimmed, "\n") {
		if IsNice(str) {
			count++
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}

func (s *Server) Solve2015day5part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-5.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	count := 0
	for _, str := range strings.Split(trimmed, "\n") {
		if IsNiceAlso(str) {
			count++
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
