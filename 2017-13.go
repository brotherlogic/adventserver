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

func computeSeverity(data string) int {
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

	for board.me <= board.mpos {
		// Update me
		board.me++

		// Check for collision
		if val, ok := board.pos[board.me]; ok && val == 0 {
			board.catches += board.me * board.lens[board.me]
		}

		// Update pos
		for key, _ := range board.pos {
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
