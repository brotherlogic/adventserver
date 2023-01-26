package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type sboard struct {
	lens    map[int]int
	pos     map[int]int
	move    map[int]bool
	me      int
	mpos    int
	catches int
}

func computeSeverityDelay(data string) int {
	start := 0
	sev := computeSeverityWithDelay(data, start)
	for sev > 0 {
		start++
		sev = computeSeverityWithDelay(data, start)
	}

	return start
}

func computeSeverity(data string) int {
	return computeSeverityWithDelay(data, 0)
}

func computeSeverityWithDelay(data string, delay int) int {
	tripZero := false
	board := sboard{lens: make(map[int]int), pos: make(map[int]int), move: make(map[int]bool), me: -1, catches: 0}

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		elems := strings.Split(line, ":")
		key := getInt32(strings.TrimSpace(elems[0]))
		value := getInt32(strings.TrimSpace(elems[1]))
		board.lens[key] = value
		board.pos[key] = 0
		board.move[key] = true

		if key > board.mpos {
			board.mpos = key
		}
	}

	timer := 0
	board.me = -1
	for board.me <= board.mpos {
		// Update me
		if timer >= delay {
			board.me++
		}

		// Check for collision
		if val, ok := board.pos[board.me]; ok && val == 0 {
			board.catches += board.me * board.lens[board.me]
			if board.me == 0 {
				tripZero = true
			}
		}

		// Update pos
		for key := range board.pos {
			if board.move[key] {
				board.pos[key]++
			} else {
				board.pos[key]--
			}

			// Wrap around
			if board.pos[key] == board.lens[key]-1 {
				board.move[key] = false
			}
			if board.pos[key] == 0 {
				board.move[key] = true
			}
		}

		timer++
	}

	// Hack adjustment since one of the catches is a zero
	if board.catches == 0 && tripZero {
		board.catches = 1
	}

	return board.catches
}

func (s *Server) Solve2017day13part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-13.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeSeverity(data))}, nil
}

func (s *Server) Solve2017day13part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-13.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeSeverityDelay(data))}, nil
}
