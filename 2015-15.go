package main

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type props struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func buildProps(data string) []props {
	var ps []props
	for _, l := range strings.Split(data, "\n") {
		bah := strings.Split(l, ":")
		elems := strings.Split(bah[1], ",")
		p := props{}
		for _, elem := range elems {
			bits := strings.Fields(elem)
			value, _ := strconv.ParseInt(bits[1], 10, 32)
			switch bits[0] {
			case "capacity":
				p.capacity = int(value)
			case "durability":
				p.durability = int(value)
			case "flavor":
				p.flavor = int(value)
			case "texture":
				p.texture = int(value)
			case "calories":
				p.calories = int(value)
			}
		}

		ps = append(ps, p)
	}

	return ps
}

func findBestIng(ap []props, maxv int) int {
	return findBestWith(make([]int, 0), ap, maxv)
}

func findBestWith(sofar []int, ap []props, maxv int) int {
	if len(sofar) == len(ap) {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0

		for i := range ap {
			capacity += sofar[i] * ap[i].capacity
			durability += sofar[i] * ap[i].durability
			flavor += sofar[i] * ap[i].flavor
			texture += sofar[i] * ap[i].texture
		}

		if sofar[0] == 44 {
			log.Printf("%v %v %v %v %+v %+v", capacity, durability, flavor, texture, ap[0], ap[1])
		}

		if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
			return 0
		}
		return capacity * durability * flavor * texture
	}

	if len(sofar) == len(ap)-1 {
		nv := maxv
		for _, c := range sofar {
			nv -= c
		}

		if nv <= 0 {
			return 0
		}

		sofar = append(sofar, nv)
		return findBestWith(sofar, ap, maxv)
	}

	best := 0
	sofar = append(sofar, 0)
	for i := 1; i < 100; i++ {
		sofar[len(sofar)-1] = i
		b := findBestWith(sofar, ap, maxv)
		if b > best {
			best = b
		}
	}

	return best
}

func computeBestScore(data string) int {

	allProps := buildProps(data)

	best := findBestIng(allProps, 100)

	return best
}

func (s *Server) Solve2015day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-15.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeBestScore(trimmed))}, nil
}
