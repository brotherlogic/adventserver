package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type octo struct {
	num     int
	flashed bool
}

func buildArr(data string) [][]*octo {
	var octoArr [][]*octo

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		var octoIn []*octo
		for _, c := range strings.TrimSpace(line) {
			num, _ := strconv.Atoi(string(c))
			octoIn = append(octoIn, &octo{num: num})
		}

		octoArr = append(octoArr, octoIn)
	}
	return octoArr
}

func restoreString(octoArr [][]*octo) string {
	data := ""
	for _, line := range octoArr {
		for _, c := range line {
			data += fmt.Sprintf("%v", c.num)
		}
		data += "\n"
	}

	return strings.TrimSpace(data)
}

func runFlash(octoArr [][]*octo) int {
	flashCount := 0

	for i, octoLine := range octoArr {
		for j, oc := range octoLine {
			if oc.num > 9 && !oc.flashed {
				oc.num = 0
				oc.flashed = true
				flashCount++

				//Run the increment
				for x := max(0, i-1); x < min(i+2, len(octoArr)); x++ {
					for y := max(0, j-1); y < min(j+2, len(octoLine)); y++ {
						if x != i || y != j {
							if !octoArr[x][y].flashed {
								octoArr[x][y].num++
							}
						}
					}
				}
			}
		}
	}

	return flashCount
}

func flash(octoArr [][]*octo) int {
	//Clear all flashes
	for _, line := range octoArr {
		for _, oc := range line {
			oc.flashed = false
		}
	}

	//First increment everything
	for _, octoLine := range octoArr {
		for _, oc := range octoLine {
			oc.num++
		}
	}

	//Now flash
	fullCount := 0
	fc := runFlash(octoArr)
	for fc != 0 {
		fullCount += fc
		fc = runFlash(octoArr)
	}

	return fullCount
}

func (s *Server) Solve2021day11part1(ctx context.Context) (*pb.SolveResponse, error) {
	data := `2566885432
	3857414357
	6761543247
	5477332114
	3731585385
	1716783173
	1277321612
	3371176148
	1162578285
	6144726367`

	arr := buildArr(data)
	count := int32(0)
	for i := 0; i < 100; i++ {
		count += int32(flash(arr))
	}

	return &pb.SolveResponse{Answer: count}, nil
}

func (s *Server) Solve2021day11part2(ctx context.Context) (*pb.SolveResponse, error) {
	data := `2566885432
	3857414357
	6761543247
	5477332114
	3731585385
	1716783173
	1277321612
	3371176148
	1162578285
	6144726367`

	arr := buildArr(data)
	count := int32(1)
	for {
		if flash(arr) == 100 {
			return &pb.SolveResponse{Answer: count}, nil
		}
		count++
	}
}
