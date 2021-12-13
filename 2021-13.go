package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runFolds(data string, numFolds int) (int, [][]bool) {
	dotCount := 0

	grid := buildGrid(data)

	folds := 0
	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "fold") && folds < numFolds {
			elems := strings.Split(line, "=")
			num, _ := strconv.Atoi(elems[1])
			if strings.Contains(line, "x") {
				grid = vertFold(grid, num)
			} else {
				grid = horizFold(grid, num)
			}
			folds++
		}
	}

	for _, line := range grid {
		for _, v := range line {
			if v {
				dotCount++
			}
		}
	}

	return dotCount, grid
}

func buildGrid(data string) [][]bool {
	var grid [][]bool
	for i := 0; i < 10000; i++ {
		grid = append(grid, make([]bool, 10000))
	}

	for _, line := range strings.Split(data, "\n") {
		tline := strings.TrimSpace(line)
		if strings.Contains(tline, ",") {
			elems := strings.Split(tline, ",")
			x, _ := strconv.Atoi(elems[0])
			y, _ := strconv.Atoi(elems[1])
			grid[y][x] = true
		}
	}

	return grid
}

func horizFold(grid [][]bool, hVal int) [][]bool {
	for i, line := range grid {
		if i > hVal {
			for j, v := range line {
				//log.Printf("%v,%v -> %v but %v", j, i, v, grid[j][hVal-(i-hVal)])
				if v && !grid[hVal-(i-hVal)][j] {
					grid[hVal-(i-hVal)][j] = true
				}
				grid[i][j] = false
			}
		}
	}
	return grid
}

func vertFold(grid [][]bool, vVal int) [][]bool {
	for i, line := range grid {
		for j, v := range line {
			if j > vVal {
				if v && !grid[i][vVal-(j-vVal)] {
					grid[i][vVal-(j-vVal)] = true
				}
				grid[i][j] = false
			}

		}
	}
	return grid
}

func printGrid(grid [][]bool) string {
	maxX := 0
	maxY := 0
	for y, line := range grid {
		for x, v := range line {
			if v {
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	res := ""
	for y, line := range grid {
		if y <= maxY {
			for x, v := range line {
				if x <= maxX {
					if v {
						res += fmt.Sprintf("#")
					} else {
						res += fmt.Sprintf(".")
					}
				}
			}

			res += fmt.Sprintf("\n")
		}
	}

	return res
}

func (s *Server) Solve2021day13part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-13.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	val, _ := runFolds(trimmed, 1)
	return &pb.SolveResponse{Answer: int32(val)}, nil
}

func (s *Server) Solve2021day13part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-13.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	_, grid := runFolds(trimmed, 900)
	return &pb.SolveResponse{StringAnswer: printGrid(grid)}, nil
}
