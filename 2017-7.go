package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type node struct {
	name     string
	value    string
	children []*node
	parent   *node
}

type tree struct {
	root *node
}

func buildTree(data string) *tree {
	var nodes []*node

	for _, line := range strings.Split(data, "\n") {
		bits := strings.Fields(line)

		nodeName := bits[0]
		value := bits[1]

		var n *node
		for _, node := range nodes {
			if node.name == nodeName {
				n = node
				n.value = value
			}
		}
		if n == nil {
			n = &node{name: nodeName, value: value}
			nodes = append(nodes, n)
		}

		if len(bits) > 2 {
			for _, child := range bits[3:] {
				if len(child) == 5 {
					child = child[0:4]
				}

				var cn *node
				for _, node := range nodes {
					if node.name == child {
						cn = node
						cn.parent = n
					}
				}

				if cn == nil {
					cn = &node{name: child, parent: n}
					nodes = append(nodes, cn)
				}
				n.children = append(n.children, cn)
			}
		}
	}

	n1 := nodes[0]
	for n1.parent != nil {
		n1 = n1.parent
	}
	return &tree{root: n1}
}

func getBottom(data string) string {
	tree := buildTree(data)

	return tree.root.name
}

func (s *Server) Solve2017day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-7.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{StringAnswer: getBottom(trimmed)}, nil
}
