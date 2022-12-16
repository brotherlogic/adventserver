package main

import (
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type sensor struct {
	sx, sy int
	bx, by int
}

func (s sensor) findBottom(x int) int {
	manDist := abs(s.bx-s.sx) + abs(s.by-s.sy)
	ydist := s.sy + (manDist - abs(s.sx-x))

	return ydist
}

func (s sensor) found(x, y int) int {
	manDist := abs(s.bx-s.sx) + abs(s.by-s.sy)
	sDist := abs(s.sx-x) + abs(s.sy-y)

	if s.bx == x && s.by == y {
		return 2
	}

	if sDist <= manDist {
		return 1
	}
	return 0
}

func parsePiece(piece string) int {
	fields := strings.Split(piece, "=")
	if strings.HasSuffix(fields[1], ",") || strings.HasSuffix(fields[1], ":") {
		return getInt32(fields[1][0 : len(fields[1])-1])
	}
	return getInt32(fields[1])
}

func parseLine(data string) (int, int, int, int) {
	fields := strings.Fields(strings.TrimSpace(data))
	return parsePiece(fields[2]), parsePiece(fields[3]), parsePiece(fields[8]), parsePiece(fields[9])
}

func maxInt(ints ...int) int {
	maxv := ints[0]
	for _, intv := range ints {
		if intv > maxv {
			maxv = intv
		}
	}

	return maxv
}

func minInt(ints ...int) int {
	minv := ints[0]
	for _, intv := range ints {
		if intv < minv {
			minv = intv
		}
	}

	return minv
}

func printSensor(arr [][]int) string {
	ret := ""
	for y := 0; y < len(arr[0]); y++ {
		for x := 0; x < len(arr); x++ {
			switch arr[x][y] {
			case 0:
				ret += "."
			case 1:
				ret += "#"
			case 2:
				ret += "S"
			case 3:
				ret += "B"
			}
		}
		ret += "\n"
	}

	return ret
}

func countKnown(data string, y int) int {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, 0, 0
	var sensors []sensor
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			sx, sy, bx, by := parseLine(line)
			dist := abs(bx-sx) + abs(by-sy)

			maxX = maxInt(maxX, sx+dist)
			minX = minInt(minX, sx-dist)
			maxY = maxInt(maxY, sy+dist)
			minY = minInt(minY, sy-dist)

			sensors = append(sensors, sensor{sx: sx, sy: sy, bx: bx, by: by})
		}
	}

	count := 0
	for x := minX; x <= maxX; x++ {
		for _, sensor := range sensors {
			if sensor.found(x, y) == 1 {
				count++
				break
			}
		}
	}

	return count
}

func findKnown(data string, mm int) int64 {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, 0, 0
	var sensors []sensor
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			sx, sy, bx, by := parseLine(line)
			dist := abs(bx-sx) + abs(by-sy)

			maxX = maxInt(maxX, sx+dist)
			minX = minInt(minX, sx-dist)
			maxY = maxInt(maxY, sy+dist)
			minY = minInt(minY, sy-dist)

			sensors = append(sensors, sensor{sx: sx, sy: sy, bx: bx, by: by})
		}
	}

	count := 0
	x := 0
	for x <= mm {
		y := 0
		for y <= mm {
			found := false
			for _, sensor := range sensors {
				if sensor.found(x, y) != 0 {
					found = true
					y = sensor.findBottom(x)
					break
				}
			}

			if !found {
				return 4000000*int64(x) + int64(y)
			}

			y++
		}
		x++
	}

	return int64(count)
}

func (s *Server) Solve2022day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-15.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countKnown(data, 2000000))}, nil
}

func (s *Server) Solve2022day15part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-15.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{BigAnswer: findKnown(data, 4000000)}, nil
}
