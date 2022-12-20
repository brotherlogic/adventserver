package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getGrid(grid [][][]int, x, y, z int) (int, bool) {
	if x >= len(grid) || x < 0 {
		return -1, false
	}

	if y >= len(grid[x]) || y < 0 {
		return -1, false
	}

	if z >= len(grid[x][y]) || z < 0 {
		return -1, false
	}

	return grid[x][y][z], true
}

func getGridExt(grid [][][]int, x, y, z int) (int, bool) {
	if x >= len(grid) || x < 0 {
		return -1, false
	}

	if y >= len(grid[x]) || y < 0 {
		return -1, false
	}

	if z >= len(grid[x][y]) || z < 0 {
		return -1, false
	}

	return grid[x][y][z], true
}

func countEdges(data string) int {
	var grid [][][]int

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		elems := strings.Split(strings.TrimSpace(line), ",")
		x, y, z := getInt32(elems[0]), getInt32(elems[1]), getInt32(elems[2])
		lg := len(grid)
		for i := 0; i <= x-lg+1; i++ {
			grid = append(grid, make([][]int, 0))
		}

		lg = len(grid[x])
		for i := 0; i <= y-lg+1; i++ {
			grid[x] = append(grid[x], make([]int, 0))
		}

		lg = len(grid[x][y])
		for i := 0; i <= z-lg+1; i++ {
			grid[x][y] = append(grid[x][y], 0)
		}

		grid[x][y][z] = 1
	}

	count := 0
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] > 0 {

					if val, ok := getGrid(grid, x+1, y, z); !ok || val == 0 {

						count++
					}

					if val, ok := getGrid(grid, x-1, y, z); !ok || val == 0 {
						count++
					}

					if val, ok := getGrid(grid, x, y+1, z); !ok || val == 0 {
						count++
					}

					if val, ok := getGrid(grid, x, y-1, z); !ok || val == 0 {
						count++
					}

					if val, ok := getGrid(grid, x, y, z+1); !ok || val == 0 {
						count++
					}

					if val, ok := getGrid(grid, x, y, z-1); !ok || val == 0 {
						count++
					}

				}
			}
		}
	}

	return count
}

func colourGrid(grid [][][]int) [][][]int {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			for z := 0; z < len(grid[x][y]); z++ {
				grid = colourGridInt(grid, x, y, z)
			}
		}
	}
	return grid
}

func colourGridInt(grid [][][]int, x, y, z int) [][][]int {
	// Base case - do nothings
	if grid[x][y][z] == 2 || grid[x][y][z] == 1 {
		return grid
	}

	if x == 0 || x == len(grid)-1 {
		grid[x][y][z] = 3
		return grid
	}

	if y == 0 || y == len(grid[x])-1 {
		grid[x][y][z] = 3
		return grid
	}

	if z == 0 || z == len(grid[x][y])-1 {
		grid[x][y][z] = 3
		return grid
	}

	// We've seen this element
	grid[x][y][z] = 1

	// x axis
	if x > 0 {
		if grid[x-1][y][z] == 0 {
			grid = colourGridInt(grid, x-1, y, z)
		}
		if grid[x-1][y][z] == 3 {
			grid[x][y][z] = 3
		}
	}
	if x < len(grid)-1 {
		if grid[x+1][y][z] == 0 {
			grid = colourGridInt(grid, x+1, y, z)
		}
		if grid[x+1][y][z] == 3 {
			grid[x][y][z] = 3
		}
	}
	if y > 0 {
		if grid[x][y-1][z] == 0 {
			grid = colourGridInt(grid, x, y-1, z)
		}
		if grid[x][y-1][z] == 3 {
			grid[x][y][z] = 3
		}
	}
	if y < len(grid[x])-1 {
		if grid[x][y+1][z] == 0 {
			grid = colourGridInt(grid, x, y+1, z)
		}
		if grid[x][y+1][z] == 3 {
			grid[x][y][z] = 3
		}
	}
	if z > 0 {
		if grid[x][y][z-1] == 0 {
			grid = colourGridInt(grid, x, y, z-1)
		}
		if grid[x][y][z-1] == 3 {
			grid[x][y][z] = 3
		}
	}
	if z < len(grid[x][y])-1 {
		if grid[x][y][z+1] == 0 {
			grid = colourGridInt(grid, x, y, z+1)
		}
		if grid[x][y][z+1] == 3 {
			grid[x][y][z] = 3
		}
	}

	return grid
}

func printLava(lava [][][]int) string {
	ret := ""

	for z := 0; z < len(lava[0][0]); z++ {
		for y := 0; y < len(lava[0]); y++ {
			for x := 0; x < len(lava); x++ {
				ret += fmt.Sprintf("%v", lava[x][y][z])
			}
			ret += "\n"
		}
		ret += "\n"
	}

	return ret
}

func countEdgesExt(data string) int {
	var grid [][][]int

	maxEdge := 20
	for x := 0; x < maxEdge; x++ {
		var iGrid [][]int
		for y := 0; y < maxEdge; y++ {
			iiGrid := make([]int, maxEdge)
			iGrid = append(iGrid, iiGrid)
		}
		grid = append(grid, iGrid)
	}

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		elems := strings.Split(strings.TrimSpace(line), ",")
		x, y, z := getInt32(elems[0]), getInt32(elems[1]), getInt32(elems[2])
		grid[x][y][z] = 2
	}

	grid = colourGrid(grid)

	count := 0
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] == 2 {

					if val, ok := getGrid(grid, x+1, y, z); !ok || val == 3 {
						count++
					}

					if val, ok := getGrid(grid, x-1, y, z); !ok || val == 3 {
						count++
					}

					if val, ok := getGrid(grid, x, y+1, z); !ok || val == 3 {
						count++
					}

					if val, ok := getGrid(grid, x, y-1, z); !ok || val == 3 {
						count++
					}

					if val, ok := getGrid(grid, x, y, z+1); !ok || val == 3 {
						count++
					}

					if val, ok := getGrid(grid, x, y, z-1); !ok || val == 3 {
						count++
					}

				}
			}
		}
	}

	return count
}

func (s *Server) Solve2022day18part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-18.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countEdges(data))}, nil
}

func (s *Server) Solve2022day18part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-18.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countEdgesExt(data))}, nil
}
