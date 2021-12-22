package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildLarge(str string) [][]bool {
	var large [][]bool
	size := 1000

	for i := 0; i < size; i++ {
		large = append(large, make([]bool, size))
	}

	elems := strings.Split(strings.TrimSpace(str), "\n")
	for i := 0; i < len(elems); i++ {
		ts := strings.TrimSpace(elems[i])
		for j := 0; j < len(ts); j++ {
			switch string(ts[j]) {
			case "#":
				large[size/2+i][size/2+j] = true
			}
		}
	}

	return large
}

func buildImageEnhance(ieh string) string {
	str := ""
	for _, line := range strings.Split(strings.TrimSpace(ieh), "\n") {
		for _, ch := range strings.TrimSpace(line) {
			str = str + string(ch)
		}
	}

	return str
}

func findBounds(image [][]bool) (int, int, int, int) {
	xmin := len(image)
	xmax := 0
	ymin := len(image[0])
	ymax := 0

	for i := range image {
		for j := range image[i] {
			if image[i][j] {
				if i < xmin {
					xmin = i
				}
				if i > xmax {
					xmax = i
				}

				if j < ymin {
					ymin = j
				}
				if j > ymax {
					ymax = j
				}
			}
		}
	}

	return xmin, xmax, ymin, ymax
}

func findRight(image [][]bool) int {
	minLeft := len(image[0])
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] && j < minLeft {
				minLeft = j
			}
		}
	}
	return minLeft
}

func doEnhance(str string, ieh string) bool {
	bval := ""
	for _, ch := range str {
		switch string(ch) {
		case "#":
			bval = bval + "1"
		case ".":
			bval = bval + "0"
		}
	}

	conv, _ := strconv.ParseInt(bval, 2, 32)
	return string(ieh[conv]) == "#"
}

func enhance(image [][]bool, ieh string) [][]bool {
	iea := buildImageEnhance(ieh)
	log.Printf("IEA: %v", iea)

	var nimage [][]bool
	for i := range image {
		nimage = append(nimage, make([]bool, len(image[i])))
	}

	xl, xr, yl, yr := findBounds(image)
	log.Printf("BOUNDS: %v,%v,%v,%v", xl, xr, yl, yr)

	for i := xl - 1; i <= xr+1; i++ {
		for j := yl - 1; j <= yr+1; j++ {
			str := ""
			log.Printf("TR: %v,%v", i, j)
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if image[x][y] {
						str = str + "#"
					} else {
						str = str + "."
					}
				}
			}

			en := doEnhance(str, iea)
			nimage[i][j] = en
		}
	}

	return nimage

}

func printImage(image [][]bool) {
	xmin, xmax, ymin, ymax := findBounds(image)

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			if image[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countLit(image [][]bool) int {
	count := 0
	for _, val := range image {
		for _, vval := range val {
			if vval {
				count++
			}
		}
	}

	return count
}

func (s *Server) Solve2021day20part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-20.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	elems := strings.Split(trimmed, "\n")
	ieh := elems[0]
	image := elems[2]
	for i := 3; i < len(elems); i++ {
		image += "\n" + elems[i]
	}

	resolve1 := enhance(buildLarge(image), ieh)
	resolve2 := enhance(resolve1, ieh)

	count := countLit(resolve2)
	return &pb.SolveResponse{Answer: int32(count)}, nil
}
