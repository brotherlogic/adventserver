package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type elfMazeNode struct {
	north, south, east, west *elfMazeNode
	x, y                     int
	wall                     bool
	name                     string
}

func wrap(val, limit int) int {
	if val%limit >= 0 {
		return val % limit
	}
	return limit + val%limit
}

func buildFunnyMaze(data string) (*elfMazeNode, []string) {
	mazeSize := 1000
	var maze [][]*elfMazeNode
	for i := 0; i < mazeSize; i++ {
		maze = append(maze, make([]*elfMazeNode, mazeSize))
	}

	for y, line := range strings.Split(data, "\n") {
		for x, char := range line {
			if char == '.' {
				maze[x][y] = &elfMazeNode{x: x + 1, y: y + 1, wall: false, name: fmt.Sprintf("%v,%v", x+1, y+1)}
			} else if char == '#' {
				maze[x][y] = &elfMazeNode{x: x + 1, y: y + 1, wall: true, name: fmt.Sprintf("%v,%v", x+1, y+1)}
			}
		}
	}

	// Fill out the maze
	for x := range maze {
		for y := range maze[x] {
			if maze[x][y] != nil {
				east := 1
				for maze[wrap(x+east, mazeSize)][y] == nil {
					east++
				}
				maze[x][y].east = maze[wrap(x+east, mazeSize)][y]

				west := 1
				for maze[wrap(x-west, mazeSize)][y] == nil {
					west++
				}
				maze[x][y].west = maze[wrap(x-west, mazeSize)][y]

				south := 1
				for maze[x][wrap(y+south, mazeSize)] == nil {
					south++
				}
				maze[x][y].south = maze[x][wrap(y+south, mazeSize)]

				north := 1
				for maze[x][wrap(y-north, mazeSize)] == nil {
					north++
				}
				maze[x][y].north = maze[x][wrap(y-north, mazeSize)]
			}
		}
	}

	var node *elfMazeNode
	for x := 0; x < mazeSize; x++ {
		if maze[x][0] != nil {
			node = maze[x][0]
			break
		}
	}

	var bits []string
	sofar := ""
	pieces := strings.Split(strings.TrimSpace(data), "\n")
	for _, char := range pieces[len(pieces)-1] {
		if char == 'R' {
			if len(sofar) > 0 {
				bits = append(bits, sofar)
				sofar = ""
			}
			bits = append(bits, "R")
		} else if char == 'L' {
			if len(sofar) > 0 {
				bits = append(bits, sofar)
				sofar = ""
			}
			bits = append(bits, "L")
		} else {
			sofar = sofar + string(char)
		}
	}
	if len(sofar) > 0 {
		bits = append(bits, sofar)
		sofar = ""
	}

	return node, bits
}

func runFunnyMaze(data string) int {
	startNode, path := buildFunnyMaze(data)

	facing := 0 // Right
	curr := startNode

	for _, node := range path {
		switch node {
		case "R":
			facing++
		case "L":
			facing--
		default:
			steps := getInt32(node)
			for i := 0; i < steps; i++ {
				switch facing {
				case 0:
					if curr.east != nil && !curr.east.wall {
						curr = curr.east
					}
				case 1:
					if curr.south != nil && !curr.south.wall {
						curr = curr.south
					}
				case 2:
					if curr.west != nil && !curr.west.wall {
						curr = curr.west
					}
				case 3:
					if curr.north != nil && !curr.north.wall {
						curr = curr.north
					}
				}
			}
		}

		if facing == -1 {
			facing = 3
		} else if facing == 4 {
			facing = 0
		}
	}

	return curr.y*1000 + 4*curr.x + facing
}

func (s *Server) Solve2022day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runFunnyMaze(data))}, nil
}
