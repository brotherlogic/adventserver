package main

import (
	"math"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type spell struct {
	name                                   string
	cost, damage, heal, armor, mana, turns int
}

func magicFight(p1, p2 player) int {
	spells := []spell{
		{"Poison", 173, 3, 0, 0, 0, 6},
		{"Magic Missile", 53, 4, 0, 0, 0, 1},
		{"Drain", 73, 2, 2, 0, 0, 1},
		{"Shield", 113, 0, 0, 7, 0, 6},
		{"Recharge", 229, 0, 0, 0, 101, 5},
	}

	return magicFightInternal(p1, p2, spells, make([]spell, 0), 0)
}

func magicFightInternal(p1, p2 player, spells, activeSpells []spell, mana int) int {
	bmana := math.MaxInt
	for _, nspell := range spells {

		if p1.mana < nspell.cost {
			continue
		}

		active := false
		for _, as := range activeSpells {
			if nspell.name == as.name && as.turns > 0 {
				active = true
			}
		}

		if active {
			continue
		}

		p1copy := p1
		p1copy.mana -= nspell.cost
		p1copy.armor += nspell.armor

		var nactiveSpells []spell
		nactiveSpells = append(nactiveSpells, activeSpells...)
		nactiveSpells = append(nactiveSpells, nspell)

		nval := magicFightRound(p1copy, p2, spells, nactiveSpells, mana+nspell.cost)

		if nval < bmana {
			bmana = nval
		}
	}

	return bmana
}

func magicFightRound(p1, p2 player, spells, activeSpells []spell, mana int) int {
	for t := 0; t < 2; t++ {
		for _, as := range activeSpells {
			if as.turns > 0 {
				p2.hitp -= as.damage
				p1.hitp += as.heal
				p1.mana += as.mana

				if p2.hitp <= 0 {
					return mana
				}

				if as.turns == 1 {
					p1.armor -= as.armor
				}

				as.turns--
			}
		}
	}

	p1.hitp -= max(1, p2.damage-p1.armor)
	if p1.hitp <= 0 {
		return math.MaxInt
	}

	return magicFightInternal(p1, p2, spells, activeSpells, mana)
}

func (s *Server) Solve2015day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(magicFight(player{hitp: 50, mana: 500}, player{hitp: 71, damage: 10}))}, nil
}
