package main

import (
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type spell struct {
	name                                   string
	cost, damage, heal, armor, mana, turns int
	cd                                     string
}

func magicFight(p1, p2 player) int {
	spells := []spell{
		{"Poison", 173, 3, 0, 0, 0, 6, "P"},
		{"Magic Missile", 53, 4, 0, 0, 0, 1, "M"},
		{"Drain", 73, 2, 2, 0, 0, 1, "D"},
		{"Shield", 113, 0, 0, 7, 0, 6, "S"},
		{"Recharge", 229, 0, 0, 0, 101, 5, "R"},
	}

	val, _ := magicFightInternal(p1, p2, spells, make([]spell, 0), 0, "")
	return val
}

func magicFightInternal(p1, p2 player, spells, activeSpells []spell, mana int, cast string) (int, string) {
	if len(cast) > 10 {
		return math.MaxInt, cast
	}
	bmana := math.MaxInt
	bcast := ""

	for _, nspell := range spells {
		if strings.HasPrefix("RSDPM", cast) {
			//log.Printf("TRYING %v %v and %v [%v]", nspell.name, p1.hitp, p2.hitp, cast)
		}

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
		ncast := cast + nspell.cd

		nval, cast := magicFightRound(p1copy, p2, spells, nactiveSpells, mana+nspell.cost, ncast)
		//log.Printf("%v -> %v", nval, cast)

		if nval < bmana {
			bmana = nval
			bcast = cast
		}
	}

	return bmana, bcast
}

func magicFightRound(p1, p2 player, spells, activeSpells []spell, mana int, cast string) (int, string) {
	for t := 0; t < 2; t++ {
		for i := range activeSpells {
			if activeSpells[i].turns > 0 {
				//log.Printf("%v -> %v : %v [%v]", cast, activeSpells[i].name, p2.hitp-activeSpells[i].damage, activeSpells[i].turns)
				p2.hitp -= activeSpells[i].damage
				p1.hitp += activeSpells[i].heal
				p1.mana += activeSpells[i].mana

				//log.Printf("HIT = %v / %v", p2.hitp, p1.hitp)
				if p2.hitp <= 0 {
					return mana, cast
				}

				if activeSpells[i].turns == 0 {
					p1.armor -= activeSpells[i].armor
				}

				activeSpells[i].turns = activeSpells[i].turns - 1
			}
		}
	}

	p1.hitp -= max(1, p2.damage-p1.armor)
	//log.Printf("HIT = %v / %v", p2.hitp, p1.hitp)
	if p1.hitp <= 0 {
		return math.MaxInt, cast
	}

	return magicFightInternal(p1, p2, spells, activeSpells, mana, cast)
}

func (s *Server) Solve2015day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(magicFight(player{hitp: 50, mana: 500}, player{hitp: 71, damage: 10}))}, nil
}