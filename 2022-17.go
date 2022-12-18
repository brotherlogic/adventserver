package main

import (
	"log"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type pixel struct {
	x, y []int
}

func getRock(count int, height int) *pixel {
	switch count % 5 {
	case 0:
		return &pixel{x: []int{2, 3, 4, 5}, y: []int{height + 3, height + 3, height + 3, height + 3}}
	case 1:
		return &pixel{x: []int{2, 3, 3, 3, 4}, y: []int{height + 4, height + 3, height + 4, height + 5, height + 4}}
	case 2:
		return &pixel{x: []int{2, 3, 4, 4}, y: []int{height + 3, height + 3, height + 3, height + 5}}
	case 3:
		return &pixel{x: []int{2, 2, 2, 2}, y: []int{height + 3, height + 4, height + 5, height + 6}}
	case 4:
		return &pixel{x: []int{2, 2, 3, 3}, y: []int{height + 3, height + 4, height + 3, height + 4}}
	default:
		log.Fatalf("No rock: %v", count)
	}

	return &pixel{}
}

func (p *pixel) move(char rune, chamber [][]int) bool {
	switch char {
	case '>':
		// Move right

		for i := range p.y {
			if p.x[i] == 6 || chamber[p.x[i]+1][p.y[i]] == 1 {
				return false
			}
		}

		for i := range p.x {
			p.x[i]++
		}

	case '<':
		for i := range p.y {
			if p.x[i] == 0 || chamber[p.x[i]-1][p.y[i]] == 1 {
				return false
			}
		}

		for i := range p.x {
			p.x[i]--
		}

	case '^':
		for i := range p.y {
			if p.y[i]-1 < 0 || chamber[p.x[i]][p.y[i]-1] == 1 {
				return false
			}
		}

		for i := range p.y {
			p.y[i]--
		}
	}

	return true
}

func runTetris(data string, maxv int) []int {
	var chamber [][]int
	rows := 10000
	for i := 0; i < 7; i++ {
		chamber = append(chamber, make([]int, rows))
	}

	count := 0
	mpointer := 0
	heights := make([]int, 7)

	for count < maxv {
		rock := getRock(count, getHeight(heights))
		for {
			rock.move(rune(data[mpointer%len(data)]), chamber)
			mpointer++
			moved := rock.move('^', chamber)
			if !moved {
				for i := range rock.x {
					chamber[rock.x[i]][rock.y[i]] = 1
					if rock.y[i]+1 > heights[rock.x[i]] {
						heights[rock.x[i]] = rock.y[i] + 1
					}
				}
				break
			}
		}
		count++
	}

	return heights
}

func getHeight(h []int) int {
	highest := 0
	for _, value := range h {
		if value > highest {
			highest = value
		}
	}

	return highest
}

func (s *Server) Solve2022day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-17.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getHeight(runTetris(data, 2022)))}, nil
}
