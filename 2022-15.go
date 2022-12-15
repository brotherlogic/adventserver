package main

import (
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	sline = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2022_15_line",
		Help: "The number of server requests",
	})
	dline = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_2022_15_down",
		Help: "The number of server requests",
	})
)

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
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			sx, sy, bx, by := parseLine(line)
			dist := abs(bx-sx) + abs(by-sy)

			maxX = maxInt(maxX, sx+dist)
			minX = minInt(minX, sx-dist)
			maxY = maxInt(maxY, sy+dist)
			minY = minInt(minY, sy-dist)
		}
	}

	var arr [][]int
	for i := minX; i <= maxX-minX; i++ {
		arr = append(arr, make([]int, maxY-minY+1))
	}

	for i, line := range strings.Split(data, "\n") {
		sline.Set(float64(i))
		if len(strings.TrimSpace(line)) > 0 {
			sx, sy, bx, by := parseLine(line)
			arr[sx-minX][sy-minY] = 2
			arr[bx-minX][by-minY] = 3
			manDist := abs(sx-bx) + abs(sy-by)

			for postx := sx - manDist; postx <= sx+manDist; postx++ {
				dline.Set(float64(sx + manDist - postx))
				for posty := sy - manDist; posty <= sy+manDist; posty++ {
					coordx, coordy := postx-minX, posty-minY
					if abs(sx-postx)+abs(sy-posty) <= manDist {
						if arr[coordx][coordy] == 0 {
							arr[coordx][coordy] = 1
						}
					}
				}
			}
		}
	}

	count := 0
	for x := 0; x < len(arr); x++ {
		if arr[x][y-minY] == 1 {
			count++
		}
	}

	return count
}

func (s *Server) Solve2022day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-15.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(countKnown(data, 2000000))}, nil
}
