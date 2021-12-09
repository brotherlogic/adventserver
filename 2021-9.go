package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func spider(matrix [][]int, x, y int, repl int) {
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || matrix[x][y] == 9 || matrix[x][y] == repl {
		return
	}

	matrix[x][y] = repl
	spider(matrix, x-1, y, repl)
	spider(matrix, x+1, y, repl)
	spider(matrix, x, y-1, repl)
	spider(matrix, x, y+1, repl)

}

func fillMatrix(matrix [][]int, x, y int, repl int) int {
	spider(matrix, x, y, repl)

	count := 0
	for _, val := range matrix {
		for _, aval := range val {
			if aval == repl {
				count++
			}
		}
	}
	log.Printf("SPIDER[%v:%v,%v]: %v -> %v", repl, x, y, matrix, count)

	return count
}

func getRisk(data string) (int, int) {
	var matrix [][]int
	var basins []int
	basinSig := 10

	for i, line := range strings.Split(data, "\n") {
		tline := strings.TrimSpace(line)
		matrix = append(matrix, make([]int, len(tline)))
		for j, char := range tline {
			val, _ := strconv.Atoi(string(char))
			matrix[i][j] = val
		}
	}

	height := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < 9 {
				lowest := true
				for x := max(i-1, 0); x <= min(i+1, len(matrix)-1); x++ {
					for y := max(j-1, 0); y <= min(j+1, len(matrix[i])-1); y++ {
						if (x == i || y == j) && !(x == i && y == j) {
							if matrix[x][y] <= matrix[i][j] {
								lowest = false
							}
						}
					}
				}

				if lowest {
					height += matrix[i][j] + 1
					count := fillMatrix(matrix, i, j, basinSig)
					basins = append(basins, count)
					basinSig++
				}
			}
		}
	}

	sort.Ints(basins)
	log.Printf("BASINS %v", basins)

	return height, basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func (s *Server) Solve2021day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-9.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	n1, _ := getRisk(trimmed)
	return &pb.SolveResponse{Answer: int32(n1)}, nil
}

func (s *Server) Solve2021day9part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-9.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	_, n2 := getRisk(trimmed)
	return &pb.SolveResponse{Answer: int32(n2)}, nil
}
