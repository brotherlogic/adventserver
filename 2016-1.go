package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

func runCoords(str string) int {
	coords := strings.Split(str, ",")
	heading := 0
	cx := 0
	cy := 0

	for _, cd := range coords {
		c := strings.TrimSpace(cd)
		char := c[0]
		switch char {
		case 'L':
			heading++
		case 'R':
			heading--
		default:
			log.Fatalf("UNPROC: %v", cd)
		}

		count, err := strconv.Atoi(c[1:len(c)])
		if err != nil {
			log.Fatalf("Bad convert: %v (%v)", err, c)
		}

		heading = heading % 4
		switch heading {
		case 0:
			cy += count
		case 1, -3:
			cx += count
		case 2, -2:
			cy -= count
		case 3, -1:
			cx -= count
		default:
			log.Fatalf("Cannot procecss heading %v", heading)
		}
	}

	return abs(cx) + abs(cy)
}

func repeat(str string) int {
	coords := strings.Split(str, ",")
	heading := 0
	cx := 0
	cy := 0
	var seens []string
	seens = append(seens, "0,0")

	for _, cd := range coords {
		c := strings.TrimSpace(cd)
		char := c[0]
		switch char {
		case 'L':
			heading++
		case 'R':
			heading--
		default:
			log.Fatalf("UNPROC: %v", cd)
		}

		count, err := strconv.Atoi(c[1:len(c)])
		if err != nil {
			log.Fatalf("Bad convert: %v (%v)", err, c)
		}

		heading = heading % 4
		switch heading {
		case 0:
			for y := 1; y <= count; y++ {
				cy++
				str := fmt.Sprintf("%v,%v", cx, cy)
				for _, seen := range seens {
					if seen == str {
						splits := strings.Split(seen, ",")
						a1, _ := strconv.Atoi(splits[0])
						a2, _ := strconv.Atoi(splits[1])
						return abs(a1) + abs(a2)
					}
				}
				seens = append(seens, str)
			}
		case 1, -3:
			for x := 1; x <= count; x++ {
				cx++
				str := fmt.Sprintf("%v,%v", cx, cy)
				for _, seen := range seens {
					if seen == str {
						splits := strings.Split(seen, ",")
						a1, _ := strconv.Atoi(splits[0])
						a2, _ := strconv.Atoi(splits[1])
						return abs(a1) + abs(a2)
					}

				}
				seens = append(seens, str)
			}
		case 2, -2:
			for y := 1; y <= count; y++ {
				cy--
				str := fmt.Sprintf("%v,%v", cx, cy)
				for _, seen := range seens {
					if seen == str {
						splits := strings.Split(seen, ",")
						a1, _ := strconv.Atoi(splits[0])
						a2, _ := strconv.Atoi(splits[1])
						return abs(a1) + abs(a2)
					}

				}
				seens = append(seens, str)
			}
		case 3, -1:
			for x := 1; x <= count; x++ {
				cx--
				str := fmt.Sprintf("%v,%v", cx, cy)
				for _, seen := range seens {
					if seen == str {
						splits := strings.Split(seen, ",")
						a1, _ := strconv.Atoi(splits[0])
						a2, _ := strconv.Atoi(splits[1])
						return abs(a1) + abs(a2)
					}

				}
				seens = append(seens, str)
			}
		default:
			log.Fatalf("Cannot procecss heading %v", heading)
		}

	}

	return -1
}

func (s *Server) solve2016day1part1(ctx context.Context) (int32, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-1.txt")
	if err != nil {
		return -1, err
	}

	return int32(runCoords(data)), nil
}

func (s *Server) solve2016day1part2(ctx context.Context) (int32, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-1.txt")
	if err != nil {
		return -1, err
	}

	return int32(repeat(data)), nil
}
