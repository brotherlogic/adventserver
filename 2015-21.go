package main

import (
	"math"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type player struct {
	hitp, armor, damage int
}

func fight(p1, p2 player) bool {

	p1hp := p1.hitp
	p2hp := p2.hitp

	for {
		p2hp -= p1.damage - p2.armor

		if p2hp <= 0 {
			return true
		}

		p1hp -= p2.damage - p1.armor

		if p1hp <= 0 {
			return false
		}
	}
}

type adjust struct {
	cost, damage, armor int
}

func runFight(ehitp, edamage, earm int) int {
	weapons := []adjust{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
	armor := []adjust{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}
	rings1 := []adjust{
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
	}
	rings2 := []adjust{
		{0, 0, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	best := math.MaxInt
	for _, w := range weapons {
		for _, a := range armor {
			for _, r1 := range rings1 {
				for _, r2 := range rings2 {
					me := player{hitp: 100,
						damage: w.damage + r1.damage,
						armor:  a.armor + r2.armor}
					if fight(me, player{hitp: ehitp, damage: edamage, armor: earm}) {
						if w.cost+a.cost+r1.cost+r2.cost < best {
							best = w.cost + a.cost + r1.cost + r2.cost
						}
					}
				}
			}
		}
	}

	return best
}

func (s *Server) Solve2015day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runFight(103, 9, 2))}, nil
}
