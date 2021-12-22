package main

import (
	"log"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runGame(p1, p2 int, maxr int) (int, int) {
	p1sc := 0
	p2sc := 0
	dice := 0
	rolls := 0

	for {
		// Player 1 go
		dice += 3
		move := dice + dice - 1 + dice - 2
		rolls += 3
		p1 += move
		p1 = ((p1 - 1) % 10) + 1
		log.Printf("Player 1 moved %v on space %v -> %v", move, p1, p1sc+p1)
		p1sc += p1

		if p1sc >= maxr {
			break
		}

		// Player 2 go
		dice += 3
		move = dice + dice - 1 + dice - 2
		rolls += 3
		p2 += move
		p2 = ((p2 - 1) % 10) + 1
		log.Printf("Player 1 moved %v on space %v -> %v", move, p2, p2sc+p2)

		p2sc += p2

		if p2sc >= maxr {
			break
		}
	}

	if p1sc >= maxr {
		return p2sc, rolls
	} else {
		return p1sc, rolls
	}
}

func (s *Server) Solve2021day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	l, r := runGame(3, 10, 1000)
	return &pb.SolveResponse{Answer: int32(l * r)}, nil
}
