package main

import (
	"fmt"
	"log"
	"strings"

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
		return &pixel{x: []int{2, 3, 4, 4, 4}, y: []int{height + 3, height + 3, height + 3, height + 4, height + 5}}
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
			if p.x[i] == 6 || chamber[p.x[i]+1][p.y[i]] > 0 {
				return false
			}
		}

		for i := range p.x {
			p.x[i]++
		}

	case '<':
		for i := range p.y {
			if p.x[i] == 0 || chamber[p.x[i]-1][p.y[i]] > 0 {
				return false
			}
		}

		for i := range p.x {
			p.x[i]--
		}

	case '^':
		for i := range p.y {
			if p.y[i]-1 < 0 || chamber[p.x[i]][p.y[i]-1] > 0 {
				return false
			}
		}

		for i := range p.y {
			p.y[i]--
		}
	}

	return true
}

func printChamber(chamber [][]int) string {
	ret := ""
	for y := len(chamber[0]) - 1; y >= 0; y-- {
		ret += fmt.Sprintf("%v:", y)
		for x := 0; x < len(chamber); x++ {
			switch chamber[x][y] {
			case 0:
				ret += "."
			case 1:
				ret += "#"
			}
		}
		ret += "\n"
	}

	return ret
}

func runTetris(data string, maxv int) ([]int, [][]int) {
	var chamber [][]int
	rows := 100000
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
					chamber[rock.x[i]][rock.y[i]] = count
					if rock.y[i]+1 > heights[rock.x[i]] {
						heights[rock.x[i]] = rock.y[i] + 1
					}
				}
				break
			}
		}
		count++

	}

	for y := range chamber[0] {
		found := true
		for x, _ := range chamber {
			if chamber[x][y] == 0 {
				found = false
			}
		}

		if found {
		}
	}

	return heights, chamber
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

func printRow(chamber [][]int, row int) string {
	ret := ""
	for x := 0; x < 7; x++ {
		switch chamber[x][row] {
		case 0:
			ret += "."
		default:
			ret += fmt.Sprintf("%v|", chamber[x][row])
		}
	}
	return ret
}

func matchTetris(chamber [][]int, row1, row2 int) bool {
	for x := 0; x < 7; x++ {
		if chamber[x][row1] == 0 && chamber[x][row2] > 0 {
			return false
		}
		if chamber[x][row1] > 0 && chamber[x][row2] == 0 {
			return false
		}
	}
	return true
}

func applyAdditionMultiplier(chamber [][]int, addition, bounce int, goal int64) ([][]int64, int64) {

	adder := int64(bounce) * (int64(goal) / int64(bounce))

	var nchamber [][]int64

	for x := 0; x < len(chamber); x++ {
		nchamber = append(nchamber, make([]int64, len(chamber[0])))
		for y := 0; y < len(chamber[0]); y++ {
			if chamber[x][y] > 0 {
				nchamber[x][y] = int64(chamber[x][y]) + adder
			}
		}
	}
	return nchamber, int64(addition) * (int64(goal) / int64(bounce))
}

func findRepeat(chamber [][]int, top int, goal int64) int64 {

	fmt.Printf("%v\n", printRow(chamber, top))

	var nchamber [][]int64
	height := int64(0)

	for y := top - 1; y >= 2; y-- {
		if matchTetris(chamber, top, y) {
			mlen := 1
			for {
				if !matchTetris(chamber, top-mlen, y-mlen) {
					break
				}
				mlen++
			}
			if mlen > 100 {
				countBounce := 0
				for x := 0; x < 7; x++ {
					if chamber[x][y] > 0 {
						countBounce = chamber[x][top] - chamber[x][y]
					}
				}
				nchamber, height = applyAdditionMultiplier(chamber, top-y, countBounce, goal)
				break
			}
		}
	}

	for y := len(nchamber[0]) - 1; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if nchamber[x][y] == goal {
				return int64(y) + height
			}
		}
	}

	return -1

}

func (s *Server) Solve2022day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-17.txt")
	if err != nil {
		return nil, err
	}

	res, _ := runTetris(strings.TrimSpace(data), 2022)
	return &pb.SolveResponse{Answer: int32(getHeight(res))}, nil
}

func (s *Server) Solve2022day17part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-17.txt")
	if err != nil {
		return nil, err
	}

	tetis := 3000
	r, chamber := runTetris(data, tetis)
	res := getHeight(r)
	rep := findRepeat(chamber, res-20, 1000000000000)
	return &pb.SolveResponse{BigAnswer: int64(rep)}, nil
}
