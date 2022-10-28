package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runDiracGame(p1, p2 int64, p1sc, p2sc int64, play bool) (int64, int64) {
	if p1sc >= 21 {
		return 1, 0
	}

	if p2sc >= 21 {
		return 0, 1
	}

	sum1 := int64(0)
	sum2 := int64(0)
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				dice := int64(d1 + d2 + d3)
				if play {
					// Player 1 takes a go
					p1 += dice
					p1 = ((p1 - 1) % 10) + 1
					w, l := runDiracGame(p1, p2, p1sc+p1, p2sc, !play)
					sum1 += w
					sum2 += l
				} else {
					// Player 1 takes a go
					p2 += dice
					p2 = ((p2 - 1) % 10) + 1
					w, l := runDiracGame(p1, p2, p1sc, p2sc+p2, !play)
					sum1 += w
					sum2 += l
				}
			}
		}
	}

	return sum1, sum2
}

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
