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

func bestTree(data string) int {
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
