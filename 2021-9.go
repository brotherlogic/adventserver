package main

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getRisk(data string) int {
	var matrix [][]int

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
			lowest := true
			for x := max(i-1, 0); x <= min(i+1, len(matrix)-1); x++ {
				for y := max(j-1, 0); y <= min(j+1, len(matrix[i])-1); y++ {
					if (x == i || y == j) && !(x == i && y == j) {
						log.Printf("%v,%v [%v] -> %v,%v [%v]", i, j, matrix[i][j], x, y, matrix[x][y])
						if matrix[x][y] <= matrix[i][j] {
							lowest = false
						}
					}
				}
			}

			//log.Printf("FFOUND")
			if lowest {
				height += matrix[i][j] + 1
			}
		}
	}

	return height
}

func (s *Server) Solve2021day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-9.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	n1 := getRisk(trimmed)
	return &pb.SolveResponse{Answer: int32(n1)}, nil
}
