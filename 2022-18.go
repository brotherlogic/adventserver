package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

var (
	gridh = 30
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

type gn struct {
	x, y, z int
}

func getNextLava(grid [][][]int, x, y, z int, seen []gn) []gn {
	var next []gn

	if x > 0 && grid[x-1][y][z] != 2 {
		next = append(next, gn{x - 1, y, z})
	}
	if y > 0 && grid[x][y-1][z] != 2 {
		next = append(next, gn{x, y - 1, z})
	}
	if z > 0 && grid[x][y][z-1] != 2 {
		next = append(next, gn{x, y, z - 1})
	}
	if x < gridh-1 && grid[x+1][y][z] != 2 {
		next = append(next, gn{x + 1, y, z})
	}
	if y < gridh-1 && grid[x][y+1][z] != 2 {
		next = append(next, gn{x, y + 1, z})
	}
	if z < gridh-1 && grid[x][y][z+1] != 2 {
		next = append(next, gn{x, y, z + 1})
	}

	var nnext []gn
	for _, n := range next {
		found := false
		for _, n2 := range seen {
			if n2.x == n.x && n2.y == n.y && n2.z == n.z {
				found = true
				break
			}
		}
		if !found {
			nnext = append(nnext, n)
		}
	}

	return nnext
}

func canEscape(grid [][][]int, x, y, z int) bool {
	queue := []gn{{x, y, z}}
	seen := []gn{{x, y, z}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.x == 0 || head.x == gridh-1 || head.y == 0 || head.y == gridh-1 || head.z == 0 || head.z == gridh-1 {
			return true
		}

		nodes := getNextLava(grid, head.x, head.y, head.z, seen)
		seen = append(seen, nodes...)
		queue = append(queue, nodes...)
	}

	return false
}

func colourGridInt(grid [][][]int, x, y, z int) [][][]int {
	if grid[x][y][z] == 2 {
		return grid
	}
	if canEscape(grid, x, y, z) {
		grid[x][y][z] = 3
	} else {
		grid[x][y][z] = 1
	}
	return grid
}

func printLava(lava [][][]int) string {
	ret := ""

	for z := 0; z < len(lava[0][0]); z++ {
		ret += fmt.Sprintf("LAYER %v\n", z)
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

	maxEdge := gridh
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
