package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildIntArr(data string) [][]int {
	var arr [][]int

	for _, line := range strings.Split(data, "\n") {
		var aline []int
		for _, c := range strings.TrimSpace(line) {
			val, _ := strconv.Atoi(string(c))
			aline = append(aline, val)
		}
		arr = append(arr, aline)
	}

	return arr
}

type cavern struct {
	arr [][]int
}

type coord struct {
	x     int
	y     int
	score int
}

func addScore(lows []*coord, coord *coord) []*coord {
	for i, val := range lows {
		if val.score > coord.score {
			lows = append(lows[:i+1], lows[i:]...)
			lows[i] = coord
			return lows
		}
	}

	return append(lows, coord)
}

func backTracePath(cav *cavern, score [][]int, lows []*coord) {
	if len(lows) == 0 {
		return
	}
	lowest := lows[0]
	lows = lows[1:]

	if cav.arr[lowest.x][lowest.y] > lowest.score {
		cav.arr[lowest.x][lowest.y] = lowest.score

		if lowest.x > 0 {
			if lowest.score+score[lowest.x-1][lowest.y] < cav.arr[lowest.x-1][lowest.y] {
				lows = addScore(lows, &coord{x: lowest.x - 1, y: lowest.y, score: lowest.score + score[lowest.x-1][lowest.y]})
			}
		}

		if lowest.x < len(score)-1 {
			if lowest.score+score[lowest.x+1][lowest.y] < cav.arr[lowest.x+1][lowest.y] {
				lows = addScore(lows, &coord{x: lowest.x + 1, y: lowest.y, score: lowest.score + score[lowest.x+1][lowest.y]})
			}
		}

		if lowest.y > 0 {
			if lowest.score+score[lowest.x][lowest.y-1] < cav.arr[lowest.x][lowest.y-1] {
				lows = addScore(lows, &coord{x: lowest.x, y: lowest.y - 1, score: lowest.score + score[lowest.x][lowest.y-1]})
			}
		}

		if lowest.y < len(score)-1 {
			if lowest.score+score[lowest.x][lowest.y+1] < cav.arr[lowest.x][lowest.y+1] {
				lows = addScore(lows, &coord{x: lowest.x, y: lowest.y + 1, score: lowest.score + score[lowest.x][lowest.y+1]})
			}
		}
	}

	backTracePath(cav, score, lows)
}

func computePath(arr [][]int) [][]int {
	var narr [][]int
	for i := 0; i < len(arr); i++ {
		var nnar []int
		for j := 0; j < len(arr); j++ {
			nnar = append(nnar, math.MaxInt32)
		}
		narr = append(narr, nnar)

	}

	cavern := &cavern{arr: narr}
	backTracePath(cavern, arr, []*coord{{
		x:     len(arr) - 1,
		y:     len(arr[0]) - 1,
		score: arr[len(arr)-1][len(arr[0])-1],
	}})

	//printCavern(cavern)

	return narr
}

func printCavern(c *cavern) {
	fmt.Printf("\n")
	for _, val := range c.arr {
		for _, nval := range val {
			fmt.Printf("%v, ", nval)
		}
		fmt.Printf("\n")
	}
}

func getBestPath(data string) int {
	arr := buildIntArr(data)

	path := computePath(arr)

	return path[0][0] - arr[0][0]
}

func (s *Server) Solve2021day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-15.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	arr := buildIntArr(trimmed)
	s.Log(fmt.Sprintf("%v vs %v", len(arr), len(arr[0])))

	path := getBestPath(trimmed)
	return &pb.SolveResponse{Answer: int32(path)}, nil
}
