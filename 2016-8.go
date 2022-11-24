package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runLightProgram(x, y int, program string) [][]bool {
	b := make([][]bool, y)
	for i := range b {
		b[i] = make([]bool, x)
	}

	for _, line := range strings.Split(program, "\n") {
		elems := strings.Fields(line)
		switch elems[0] {
		case "rect":
			bits := strings.Split(elems[1], "x")
			val1, _ := strconv.ParseInt(bits[0], 10, 32)
			val2, _ := strconv.ParseInt(bits[1], 10, 32)

			for i := 0; i < int(val1); i++ {
				for j := 0; j < int(val2); j++ {
					b[j][i] = true
				}
			}
		case "rotate":
			if elems[1] == "column" {
				bits := strings.Split(elems[2], "=")
				num, _ := strconv.ParseInt(bits[1], 10, 32)
				by, _ := strconv.ParseInt(elems[4], 10, 32)

				ncol := make([]bool, y)

				for i := 0; i < y; i++ {
					nindex := i - int(by)
					if nindex < 0 {
						nindex += y
					} else if nindex > y {
						nindex -= y
					}
					ncol[i] = b[nindex][num]
				}

				for i := 0; i < y; i++ {
					b[i][num] = ncol[i]
				}
			}
			if elems[1] == "row" {
				bits := strings.Split(elems[2], "=")
				num, _ := strconv.ParseInt(bits[1], 10, 32)
				by, _ := strconv.ParseInt(elems[4], 10, 32)

				ncol := make([]bool, x)

				for i := 0; i < x; i++ {
					nindex := i - int(by)
					if nindex < 0 {
						nindex += x
					} else if nindex > x {
						nindex -= x
					}
					ncol[i] = b[num][nindex]
				}

				for i := 0; i < x; i++ {
					b[num][i] = ncol[i]
				}
			}
		}
	}

	return b
}

func countBoolArr(b [][]bool) int {
	count := 0
	for x := range b {
		for y := range b[x] {
			if b[x][y] {
				count++
			}
		}
	}
	return count
}

func (s *Server) Solve2016day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-8.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countBoolArr(runLightProgram(50, 6, data)))}, nil
}
