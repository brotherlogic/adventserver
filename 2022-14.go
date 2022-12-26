package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getInt32(s string) int {
	val, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	if err != nil {
		fmt.Printf("Bad int: '%v'", s)
		log.Fatalf("Bad int: '%v'", s)
	}

	return int(val)
}

func printArr(arr [][]int) string {
	ret := ""
	for y := 0; y < len(arr[0]); y++ {
		for x := 0; x < len(arr); x++ {
			switch arr[x][y] {
			case 0:
				ret += "."
			case 1:
				ret += "#"
			case 2:
				ret += "o"
			}
		}
		ret += "\n"
	}

	return ret
}

func buildSand(s string, left, right, bottom int) [][]int {
	var arr [][]int

	for x := 0; x <= right-left; x++ {
		line := make([]int, bottom+1)
		arr = append(arr, line)
	}

	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "->") {
			elems := strings.Split(line, "->")
			for i := range elems[1:] {
				parts := strings.Split(elems[i], ",")
				valxs := getInt32(parts[0])
				valys := getInt32(parts[1])
				parte := strings.Split(elems[i+1], ",")
				valxe := getInt32(parte[0])
				valye := getInt32(parte[1])

				if valxs > valxe {
					valxs, valxe = valxe, valxs
				}

				if valys > valye {
					valys, valye = valye, valys
				}

				for x := valxs; x <= valxe; x++ {
					for y := valys; y <= valye; y++ {
						arr[x-left][y] = 1
					}
				}
			}
		}
	}

	return arr
}

func findEdge(data string) (int, int, int) {
	left := math.MaxInt
	right := 0
	bottom := 0
	for _, line := range strings.Split(data, "\n") {
		for _, elem := range strings.Split(line, "->") {
			parts := strings.Split(elem, ",")
			if len(parts) == 2 {
				valx := getInt32(parts[0])
				valy := getInt32(parts[1])

				if valx < left {
					left = valx
				}
				if valx > right {
					right = valx
				}
				if valy > bottom {
					bottom = valy
				}
			}
		}
	}

	return left, right, bottom
}

func runSand(arr [][]int, left int) (int, int) {
	sandx, sandy := 500-left, 0

	for sandy < len(arr[0]) {
		if arr[sandx][sandy+1] == 0 {
			sandy++
		} else {
			if sandx <= 0 {
				return -1, -1
			}
			if arr[sandx-1][sandy+1] != 0 && sandx >= len(arr)-1 {
				return -1, -1
			}
			if arr[sandx-1][sandy+1] == 0 {
				sandy++
				sandx--
			} else if arr[sandx+1][sandy+1] == 0 {
				sandy++
				sandx++
			} else {
				return sandx, sandy
			}
		}
	}

	return -1, -1
}

func countSand(data string) int {
	left, right, bottom := findEdge(data)
	sands := buildSand(data, left, right, bottom)

	count := 0
	for {
		nsandx, nsandy := runSand(sands, left)
		if nsandx == -1 {
			return count
		}
		sands[nsandx][nsandy] = 2
		count++
	}
}

func countSandBased(data string) int {
	left, right, bottom := findEdge(data)
	adjust := 10000
	sands := buildSand(data, left-adjust, right+adjust, bottom+2)

	for x := 0; x < len(sands); x++ {
		sands[x][bottom+2] = 1
	}

	count := 0
	for {
		nsandx, nsandy := runSand(sands, left-adjust)
		if sands[500-left+adjust][0] == 2 {
			return count
		}
		sands[nsandx][nsandy] = 2
		count++

	}
}

func (s *Server) Solve2022day14part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-14.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countSand(data))}, nil
}

func (s *Server) Solve2022day14part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-14.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countSandBased(data))}, nil
}
