package main

import (
	"fmt"
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type mapNode struct {
	x, y  int
	steps int
	route string
}

func buildMap(data string) ([][]int, int, int, int, int) {
	var hMap [][]int
	x, y, gx, gy := 0, 0, 0, 0

	for i, line := range strings.Split(data, "\n") {
		var hLine []int
		if len(strings.TrimSpace(line)) > 0 {
			for j, char := range line {
				if char == 'S' {
					x, y = i, j
					hLine = append(hLine, int('a'))
				} else if char == 'E' {
					gx, gy = i, j
					hLine = append(hLine, int('z'))
				} else {
					hLine = append(hLine, int(char))
				}
			}
			hMap = append(hMap, hLine)
		}
	}

	return hMap, x, y, gx, gy
}

func getNext(node mapNode, hMap [][]int) []mapNode {
	var response []mapNode

	if node.x > 0 && hMap[node.x-1][node.y] <= hMap[node.x][node.y]+1 {
		response = append(response, mapNode{x: node.x - 1, y: node.y, steps: node.steps + 1, route: node.route + fmt.Sprintf(";%v-%v", node.x-1, node.y)})
	}

	if node.x < len(hMap)-1 && hMap[node.x+1][node.y] <= hMap[node.x][node.y]+1 {
		response = append(response, mapNode{x: node.x + 1, y: node.y, steps: node.steps + 1, route: node.route + fmt.Sprintf(";%v-%v", node.x+1, node.y)})
	}

	if node.y > 0 && hMap[node.x][node.y-1] <= hMap[node.x][node.y]+1 {
		response = append(response, mapNode{x: node.x, y: node.y - 1, steps: node.steps + 1, route: node.route + fmt.Sprintf(";%v-%v", node.x, node.y-1)})
	}

	if node.y < len(hMap[0])-1 && hMap[node.x][node.y+1] <= hMap[node.x][node.y]+1 {
		response = append(response, mapNode{x: node.x, y: node.y + 1, steps: node.steps + 1, route: node.route + fmt.Sprintf(";%v-%v", node.x, node.y+1)})
	}

	rstring := ""
	for _, r := range response {
		rstring += fmt.Sprintf("%v;", string([]byte{byte(hMap[r.x][r.y])}))
	}

	return response

}

func printMap(hmap [][]int) {
	for x := range hmap {
		for y := range hmap[x] {
			fmt.Printf("%v", string([]byte{byte(hmap[x][y])}))
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%v\n", string([]byte{byte(hmap[0][1])}))
}

func buildData(data string) []string {
	return []string{data}
}

func runMultiMap(data string) int {
	best := math.MaxInt
	for _, mData := range buildData(data) {
		val, _ := runMap(mData)
		if val < best {
			val = best
		}
	}

	return best
}

func runMap(data string) (int, string) {
	hMap, x, y, gx, gy := buildMap(data)
	seen := make(map[string]bool)

	queue := []mapNode{{x: x, y: y, steps: 0, route: fmt.Sprintf("%v-%v", x, y)}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.x == gx && head.y == gy {
			return head.steps, head.route
		}

		next := getNext(head, hMap)
		for _, n := range next {
			if _, ok := seen[fmt.Sprintf("%v-%v", n.x, n.y)]; !ok {
				queue = append(queue, n)
				seen[fmt.Sprintf("%v-%v", n.x, n.y)] = true
			}
		}
	}

	return -1, ""
}

func (s *Server) Solve2022day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-12.txt")
	if err != nil {
		return nil, err
	}

	bmap, _, _, _, _ := buildMap(data)
	s.CtxLog(ctx, fmt.Sprintf("%v, %v", len(bmap), len(bmap[0])))
	res, _ := runMap(data)
	return &pb.SolveResponse{Answer: int32(res)}, nil
}
