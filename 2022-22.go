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
	nadj, sadj, eadj, wadj   int
}

func wrap(val, limit int) int {
	if val%limit >= 0 {
		return val % limit
	}
	return limit + val%limit
}

func buildFunnyMaze(data string) (*elfMazeNode, []string, [][]*elfMazeNode) {
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

	return node, bits, maze
}

func runFunnyMaze(data string) int {
	startNode, path, _ := buildFunnyMaze(data)

	return runElfMaze(startNode, path)
}

func printElf(node *elfMazeNode) string {
	return fmt.Sprintf("%v,%v E:%v,%v;%v S:%v,%v;%v W:%v,%v;%v N:%v,%v;%v",
		node.x-1, node.y-1,
		node.east.x-1, node.east.y-1, node.east.nadj,
		node.south.x-1, node.south.y-1, node.south.nadj,
		node.west.x-1, node.west.y-1, node.west.nadj,
		node.north.x-1, node.north.y-1, node.north.nadj)
}

func runElfMaze(startNode *elfMazeNode, path []string) int {
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

						if curr.eadj > 0 {
							facing = curr.eadj
						}
						if curr.eadj == -1 {
							facing = 0
						}
						curr = curr.east
					}
				case 1:
					if curr.south != nil && !curr.south.wall {

						if curr.sadj > 0 {
							facing = curr.sadj
						}
						if curr.sadj == -1 {
							facing = 0
						}
						curr = curr.south
					}
				case 2:
					if curr.west != nil && !curr.west.wall {

						if curr.wadj > 0 {
							facing = curr.wadj
						}
						if curr.wadj == -1 {
							facing = 0
						}
						curr = curr.west
					}
				case 3:
					if curr.north != nil && !curr.north.wall {

						if curr.nadj > 0 {
							facing = curr.nadj
						}
						if curr.nadj == -1 {
							facing = 0
						}
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

func runFunnyCube(data string, projection int) int {
	startNode, path, mnodes := buildFunnyMaze(data)

	// Deal with the edges
	edgeWidth := 1
	curr := startNode.east
	for curr != startNode {
		edgeWidth++
		curr = curr.east
	}

	if projection == 2 {
		edgeWidth = edgeWidth / 2
	}

	if projection == 1 {
		// 1-4 / 4-1 edge ** stays the same **
		// 1-3 / 3-1 edge
		for y := 0; y < edgeWidth; y++ {
			cx := edgeWidth * 2
			cy := y
			nx := edgeWidth + y
			ny := edgeWidth

			mnodes[cx][cy].west = mnodes[nx][ny]
			mnodes[cx][cy].wadj = 1

			mnodes[nx][ny].north = mnodes[cx][cy]
			mnodes[nx][ny].nadj = 0
		}
		// 1-6 / 6-1 edge
		for x := 0; x < edgeWidth; x++ {
			cx := edgeWidth*3 - 1
			cy := x
			nx := edgeWidth*3 + x
			ny := 3*edgeWidth - 1

			mnodes[cx][cy].east = mnodes[nx][ny]
			mnodes[cx][cy].eadj = 2

			mnodes[nx][ny].west = mnodes[cx][cy]
			mnodes[nx][ny].wadj = 0
		}
		// 1-2 / 2-1 edge
		for x := edgeWidth * 2; x < edgeWidth*3; x++ {
			cx := x
			cy := 0
			nx := edgeWidth*3 - 1 - x
			ny := edgeWidth

			mnodes[cx][cy].north = mnodes[nx][ny]
			mnodes[cx][cy].nadj = 1

			mnodes[nx][ny].north = mnodes[cx][cy]
			mnodes[nx][ny].nadj = 1
		}

		// 2-3 / 3-2 edge ** Nothing changes **
		// 2-6 / 6-2 edge
		for y := edgeWidth; y < edgeWidth*2; y++ {
			cx := 0
			cy := y
			nx := edgeWidth*5 - 1 - y
			ny := edgeWidth*3 - 1
			mnodes[cx][cy].south = mnodes[nx][ny]
			mnodes[cx][cy].sadj = 3

			mnodes[nx][ny].north = mnodes[cx][cy]
			mnodes[nx][ny].nadj = 1
		}
		// 2-1 / 1-2 edge ** Already done **
		// 2-5 / 5-2 edge
		for x := 0; x < edgeWidth; x++ {
			cx := x
			cy := edgeWidth*2 - 1
			nx := edgeWidth*3 - 1 - x
			ny := 2*edgeWidth + 3
			mnodes[cx][cy].south = mnodes[nx][ny]
			mnodes[cx][cy].sadj = 3

			mnodes[nx][ny].south = mnodes[cx][cy]
			mnodes[nx][ny].sadj = 3
		}

		// 3-4 / 4-3 edge ** Not Needed **
		// 3-2 / 2-3 edge ** Already done **
		// 1-3 / 3-1 edge ** Already done **
		// 3-5 / 5-3 edge
		for x := edgeWidth; x < edgeWidth*2; x++ {
			cx := x
			cy := edgeWidth*2 - 1
			nx := 2 * edgeWidth
			ny := edgeWidth*4 - 1 - x
			mnodes[cx][cy].south = mnodes[nx][ny]
			mnodes[cx][cy].sadj = 0

			mnodes[nx][ny].east = mnodes[cx][cy]
			mnodes[nx][ny].eadj = 3
		}

		// 4-1 / 1-4 edge ** Already done **
		// 4-5 / 5-4 edge ** Taken care of
		// 4-3 / 3-4 edge ** Already done **
		// 4-6 / 6-4 edge
		for y := edgeWidth; y < edgeWidth*2; y++ {
			cx := edgeWidth*3 - 1
			cy := y
			nx := 5*edgeWidth - y - 1
			ny := 2 * edgeWidth

			mnodes[cx][cy].east = mnodes[nx][ny]
			mnodes[cx][cy].eadj = 1

			mnodes[nx][ny].north = mnodes[cx][cy]
			mnodes[nx][ny].nadj = 2
		}

		// 5-4 / 4-5 edge ** Already done **
		// 5-6 / 6-5 edge ** Not required **
		// 5-4 / 3-5 edge ** Already done **
		// 5-2 / 2-5 edge ** Already done **

		// 6-5 / 5-6 edge ** Already done **
		// 6-4 / 4-6 edge ** Already done **
		// 6-2 / 2-6 edge ** Already done **
		// 6-1 / 1-6 edge ** Already done **
	}

	if projection == 2 {
		// 1-4 / 4-1 edge
		for y := 0; y < edgeWidth; y++ {
			cx := edgeWidth
			cy := y
			nx := 0
			ny := edgeWidth*3 - 1 - y

			mnodes[cx][cy].west = mnodes[nx][ny]
			mnodes[cx][cy].wadj = -1

			mnodes[nx][ny].west = mnodes[cx][cy]
			mnodes[nx][ny].wadj = -1
		}
		// 1-3 / 3-1 edge

		// 1-6 / 6-1 edge
		for x := edgeWidth; x < edgeWidth*2; x++ {
			cx := x
			cy := 0
			nx := 0
			ny := 2*edgeWidth + x

			mnodes[cx][cy].north = mnodes[nx][ny]
			mnodes[cx][cy].nadj = -1

			mnodes[nx][ny].west = mnodes[cx][cy]
			mnodes[nx][ny].wadj = 1
		}
		// 1-2 / 2-1 edge

		// 2-3 / 3-2 edge
		for x := edgeWidth * 2; x < edgeWidth*3; x++ {
			cx := x
			cy := edgeWidth - 1
			nx := edgeWidth*2 - 1
			ny := x - edgeWidth

			mnodes[cx][cy].south = mnodes[nx][ny]
			mnodes[cx][cy].sadj = 2

			mnodes[nx][ny].east = mnodes[cx][cy]
			mnodes[nx][ny].eadj = 3
		}
		// 2-6 / 6-2 edge
		for x := edgeWidth * 2; x < edgeWidth*3; x++ {
			cx := x
			cy := 0
			nx := x - edgeWidth*2
			ny := edgeWidth*4 - 1

			mnodes[cx][cy].north = mnodes[nx][ny]
			mnodes[cx][cy].nadj = 3

			mnodes[nx][ny].south = mnodes[cx][cy]
			mnodes[nx][ny].sadj = 1
		}
		// 2-1 / 1-2 edge ** Already done **
		// 2-5 / 5-2 edge
		for y := 0; y < edgeWidth; y++ {
			cx := edgeWidth*3 - 1
			cy := y
			nx := edgeWidth*2 - 1
			ny := edgeWidth*3 - 1 - y

			mnodes[cx][cy].east = mnodes[nx][ny]
			mnodes[cx][cy].eadj = 2

			mnodes[nx][ny].east = mnodes[cx][cy]
			mnodes[nx][ny].eadj = 2
		}

		// 3-4 / 4-3 edge ** Not Needed **
		for y := edgeWidth; y < edgeWidth*2; y++ {
			cx := edgeWidth
			cy := y
			nx := y - edgeWidth
			ny := 2 * edgeWidth

			mnodes[cx][cy].west = mnodes[nx][ny]
			mnodes[cx][cy].wadj = 1

			mnodes[nx][ny].north = mnodes[cx][cy]
			mnodes[nx][ny].nadj = -1
		}
		// 3-2 / 2-3 edge ** Already done **
		// 1-3 / 3-1 edge ** Already done **
		// 3-5 / 5-3 edge

		// 4-1 / 1-4 edge ** Already done **
		// 4-5 / 5-4 edge ** Taken care of
		// 4-3 / 3-4 edge ** Already done **
		// 4-6 / 6-4 edge

		// 5-4 / 4-5 edge ** Already done **
		// 5-6 / 6-5 edge ** Not required **
		for x := edgeWidth; x < edgeWidth*2; x++ {
			cx := x
			cy := 3*edgeWidth - 1
			nx := edgeWidth - 1
			ny := 2*edgeWidth + x

			mnodes[cx][cy].south = mnodes[nx][ny]
			mnodes[cx][cy].sadj = 2

			mnodes[nx][ny].east = mnodes[cx][cy]
			mnodes[nx][ny].eadj = 3
		}
		// 5-4 / 3-5 edge ** Already done **
		// 5-2 / 2-5 edge ** Already done **

		// 6-5 / 5-6 edge ** Already done **
		// 6-4 / 4-6 edge ** Already done **
		// 6-2 / 2-6 edge ** Already done **
		// 6-1 / 1-6 edge ** Already done **
	}

	return runElfMaze(startNode, path)

}

func (s *Server) Solve2022day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runFunnyMaze(data))}, nil
}

func (s *Server) Solve2022day22part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runFunnyCube(data, 2))}, nil
}
