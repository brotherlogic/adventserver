package main

import (
	"math"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func shotIn(x, y, lx, hx, ly, hy int, right, down bool) int {
	if x >= lx && x <= hx && y >= ly && y <= hy {
		return 0
	}

	// If we're below the target and going down, we'll never reach it
	if y < ly && down {
		return 1
	}

	// If we're going right and past the target we'll never reach it
	if x >= hx && right {
		return 1
	}

	// If we're going left and past the target we'll never reach it
	if x <= lx && !right {
		return 1
	}

	return -1
}

func throwIn(x, y int, lowx, lowy, highx, highy int) int {
	currx := 0
	xvel := x
	curry := 0
	yvel := y
	tbh := 0

	for step := 0; step < math.MaxInt16; step++ {
		//log.Printf("%v, %v @ %v,%v", currx, curry, xvel, yvel)
		if shotIn(currx, curry, lowx, highx, lowy, highy, xvel > 0, yvel < 0) == 0 {
			return tbh
		} else if shotIn(currx, curry, lowx, highx, lowy, highy, xvel > 0, yvel < 0) > 0 {
			break
		}

		currx += xvel
		curry += yvel

		if curry > tbh {
			tbh = curry
		}

		if xvel > 0 {
			xvel--
		} else if xvel < 0 {
			xvel++
		}
		yvel--
	}

	return 0
}

func findBest(lowx, highx, lowy, highy int) (int, int, int) {
	bestx := 0
	besty := 0
	bestheight := 0
	for x := 1; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			height := throwIn(x, y, lowx, lowy, highx, highy)
			if height > bestheight {
				bestx = x
				besty = y
				bestheight = height
			}
		}
	}

	return bestx, besty, bestheight
}

func (s *Server) Solve2021day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	_, _, b := findBest(111, 161, -154, -101)
	return &pb.SolveResponse{Answer: int32(b)}, nil
}
