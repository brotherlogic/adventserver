package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type df struct {
	name                   string
	size, used, avail, use int
	x, y                   int
}

func convertNum(str string) int {
	num, _ := strconv.ParseInt(str[:len(str)-1], 10, 32)
	return int(num)
}

func buildNodes(data string) []*df {
	var nodes []*df
	for _, line := range strings.Split(data, "\n") {
		if len(line) > 0 && line[0] == '/' {
			fields := strings.Fields(line)

			elems := strings.Split(fields[0], "-")
			x, _ := strconv.ParseInt(elems[1][1:], 10, 32)
			y, _ := strconv.ParseInt(elems[2][1:], 10, 32)

			df := &df{name: fields[0], x: int(x), y: int(y)}
			df.size = convertNum(fields[1])
			df.used = convertNum(fields[2])
			df.avail = convertNum(fields[3])
			df.use = convertNum(fields[4])

			nodes = append(nodes, df)
		}
	}

	return nodes
}

func calcNodes(ctx context.Context, nodes []*df, log func(context.Context, string)) int {
	count := 0
	for i, n1 := range nodes {
		for j, n2 := range nodes {
			if i != j {
				if n1.used != 0 {
					if n1.used < n2.avail {
						log(ctx, fmt.Sprintf("%+v and %+v", n1, n2))
						count++
					}
				}
			}
		}
	}

	return count
}

func runDiskMaze(data string) int {
	nodes := buildNodes(data)
	maxX, maxY := 0, 0
	emptyX, emptyY := 0, 0
	maxV := 0
	for _, node := range nodes {
		if node.x > maxX {
			maxX = node.x
		}

		if node.y > maxY {
			maxY = node.y
		}

		if node.used == 0 {
			emptyX = node.x
			emptyY = node.y
			maxV = node.avail
		}
	}

	maze := make([][]int, maxX+1)
	for i := 0; i < maxX+1; i++ {
		maze[i] = make([]int, maxY+1)
	}
	maze[emptyX][emptyY] = 1

	for _, node := range nodes {
		if node.used > maxV {
			maze[node.x][node.y] = 2
		}
	}

	maze[maxX][0] = 3

	return solveMaze(maze)
}

type mazeNode struct {
	maze  [][]int
	steps int
}

func solveMaze(maze [][]int) int {
	queue := []mazeNode{mazeNode{maze: maze, steps: 0}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.maze[0][0] == 3 {
			return head.steps
		}

		x3, y3 := 0, 0
		x1, y1 := 0, 0
		for x := range head.maze {
			for y := range head.maze[x] {
				if head.maze[x][y] == 3 {
					x3, y3 = x, y
				}
				if head.maze[x][y] == 1 {
					x1, y1 = x, y
				}
			}
		}

		for _, move := range makeMoves(head.maze, x3, y3, true) {
			nmaze := maze
			nmaze[x3][y3] = 1
			nmaze[move[0]][move[1]] = 3
			queue = append(queue, mazeNode{maze: nmaze, steps: head.steps + 1})
		}

		for _, move := range makeMoves(head.maze, x1, y1, false) {
			nmaze := maze
			nmaze[x3][y3] = 0
			nmaze[move[0]][move[1]] = 3
			queue = append(queue, mazeNode{maze: nmaze, steps: head.steps + 1})
		}
	}

	return -1
}

func makeMoves(maze [][]int, x, y int, isGoal bool) [][]int {

	if isGoal {

	}
}

func printMaze(maze [][]int) string {
	res := ""
	for y := 0; y < len(maze[0]); y++ {
		for x := 0; x < len(maze); x++ {
			res += fmt.Sprintf("%v", maze[x][y])
		}
		res += fmt.Sprintln()
	}
	return res
}

func (s *Server) Solve2016day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(calcNodes(ctx, buildNodes(data), s.CtxLog))}, nil
}

func (s *Server) Solve2016day22part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runDiskMaze(data))}, nil
}
