package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildMaze(data string) ([][]int, int, int, int) {
	var maze [][]int
	sx, sy := 0, 0
	maxv := 0

	for y, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			var cMaze []int
			for _, char := range strings.TrimSpace(line) {
				switch char {
				case '#':
					cMaze = append(cMaze, -1)
				case '.':
					cMaze = append(cMaze, 0)
				default:
					val := getInt32(string(char))
					if val == 0 {
						sy, sx = len(cMaze), y
					}
					cMaze = append(cMaze, val)
					if val > maxv {
						maxv = val
					}
				}
			}

			maze = append(maze, cMaze)
		}
	}

	return maze, sx, sy, maxv
}

type ductMazeNode struct {
	x, y  int
	seen  map[int]bool
	count int
	path  string
}

func (d ductMazeNode) rep() string {
	return fmt.Sprintf("%v-%v-%v", d.x, d.y, d.seen)
}

func printDuctMaze(maze [][]int, xs, ys int) string {
	ret := ""
	for x := 0; x < len(maze); x++ {
		for y := 0; y < len(maze[0]); y++ {
			if x == xs && y == ys {
				ret += "S"
			} else {
				switch maze[x][y] {
				case -1:
					ret += "#"
				case 0:
					ret += "."
				default:
					ret += fmt.Sprintf("%v", maze[x][y])
				}
			}
		}
		ret += "\n"
	}
	return ret
}

func mCopy(s map[int]bool) map[int]bool {
	nmap := make(map[int]bool)

	for k, v := range s {
		nmap[k] = v
	}

	return nmap
}

func generateNext(head *ductMazeNode, maze [][]int) []*ductMazeNode {
	var resp []*ductMazeNode

	if head.x < len(maze)-1 && maze[head.x+1][head.y] >= 0 {
		resp = append(resp, &ductMazeNode{x: head.x + 1, y: head.y, count: head.count + 1, seen: mCopy(head.seen),
			path: head.path + fmt.Sprintf(";%v-%v", head.x+1, head.y)})
	}

	if head.x > 0 && maze[head.x-1][head.y] >= 0 {
		resp = append(resp, &ductMazeNode{x: head.x - 1, y: head.y, count: head.count + 1, seen: mCopy(head.seen),
			path: head.path + fmt.Sprintf(";%v-%v", head.x-1, head.y)})
	}

	if head.y < len(maze[0])-1 && maze[head.x][head.y+1] >= 0 {
		resp = append(resp, &ductMazeNode{x: head.x, y: head.y + 1, count: head.count + 1, seen: mCopy(head.seen),
			path: head.path + fmt.Sprintf(";%v-%v", head.x, head.y+1)})
	}

	if head.y > 0 && maze[head.x][head.y-1] >= 0 {
		resp = append(resp, &ductMazeNode{x: head.x, y: head.y - 1, count: head.count + 1, seen: mCopy(head.seen),
			path: head.path + fmt.Sprintf(";%v-%v", head.x, head.y-1)})
	}

	return resp
}

func runDuctMaze(data string, ret bool) int {
	maze, sx, sy, maxv := buildMaze(data)

	seen := make(map[string]bool)
	queue := []*ductMazeNode{&ductMazeNode{x: sx, y: sy, seen: make(map[int]bool), path: fmt.Sprintf("%v-%v", sx, sy)}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if len(head.seen) == maxv {
			if ret && head.x == sx && head.y == sy {
				return head.count
			} else if head.x == sx && head.y == sy {
				return head.count - 1
			}
		}

		if maze[head.x][head.y] > 0 {
			head.seen[maze[head.x][head.y]] = true
		}

		nexts := generateNext(head, maze)
		for _, next := range nexts {
			if val := seen[next.rep()]; !val {
				queue = append(queue, next)
				seen[next.rep()] = true
			}
		}
	}

	return -1
}

func (s *Server) Solve2016day24part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-24.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runDuctMaze(data, false))}, nil
}

func (s *Server) Solve2016day24part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-24.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runDuctMaze(data, true))}, nil
}
