package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildTreeGrid(data string) [][]int {
	var grid [][]int
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			var internal []int
			for _, char := range line {
				val, _ := strconv.ParseInt(string(char), 10, 32)
				internal = append(internal, int(val))
			}

			grid = append(grid, internal)
		}
	}

	return grid
}

func countSeen(grid [][]int, y, x int) int {
	seenU := 0
	for ny := y - 1; ny >= 0; ny-- {
		if grid[ny][x] < grid[y][x] {
			seenU++
		} else {
			seenU++
			break
		}
	}
	seenD := 0
	for ny := y + 1; ny < len(grid); ny++ {
		if grid[ny][x] < grid[y][x] {
			seenD++
		} else {
			seenD++
			break
		}
	}
	seenL := 0
	for nx := x - 1; nx >= 0; nx-- {
		if grid[y][nx] < grid[y][x] {
			seenL++
		} else {
			seenL++
			break
		}
	}
	seenR := 0
	for nx := x + 1; nx < len(grid[0]); nx++ {
		if grid[y][nx] < grid[y][x] {
			seenR++
		} else {
			seenR++
			break
		}
	}

	return seenU * seenD * seenL * seenR
}

func bestTree(data string) int {
	grid := buildTreeGrid(data)
	var see [][]int
	for i := 0; i < len(grid); i++ {
		var s []int
		for j := 0; j < len(grid[i]); j++ {
			s = append(s, 0)
		}
		see = append(see, s)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			see[y][x] = countSeen(grid, y, x)
		}
	}

	count := 0
	for i := 0; i < len(see); i++ {
		for j := 0; j < len(see[i]); j++ {
			if see[i][j] > count {
				count = see[i][j]
			}
		}
	}

	return count
}

func countVisibleTrees(data string) int {
	grid := buildTreeGrid(data)
	var visible [][]bool
	for i := 0; i < len(grid); i++ {
		var vis []bool
		for j := 0; j < len(grid[i]); j++ {
			vis = append(vis, false)
		}
		visible = append(visible, vis)
	}

	for y := 0; y < len(grid); y++ {
		sofar := -1
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] > sofar {
				visible[y][x] = true
				sofar = grid[y][x]
			}
		}

		sofar = -1
		for x := len(grid[y]) - 1; x > 0; x-- {
			if grid[y][x] > sofar {
				visible[y][x] = true
				sofar = grid[y][x]
			}
		}
	}

	for x := 0; x < len(grid[0]); x++ {
		sofar := -1
		for y := 0; y < len(grid); y++ {
			if grid[y][x] > sofar {
				visible[y][x] = true
				sofar = grid[y][x]
			}
		}

		sofar = -1
		for y := len(grid) - 1; y > 0; y-- {
			if grid[y][x] > sofar {
				visible[y][x] = true
				sofar = grid[y][x]
			}
		}
	}

	count := 0
	for i := 0; i < len(visible); i++ {
		for j := 0; j < len(visible[i]); j++ {
			if visible[i][j] {
				count++
			}
		}
	}

	return count
}

func (s *Server) Solve2022day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-8.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countVisibleTrees(data))}, nil
}

func (s *Server) Solve2022day8part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-8.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(bestTree(data))}, nil
}
