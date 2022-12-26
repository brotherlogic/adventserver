package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type node struct {
	name     string
	value    int
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
		value = value[1 : len(value)-1]

		var n *node
		for _, node := range nodes {
			if node.name == nodeName {
				n = node
				n.value = getInt32(value)
			}
		}
		if n == nil {
			n = &node{name: nodeName, value: getInt32(value)}
			nodes = append(nodes, n)
		}

		if len(bits) > 2 {
			for _, child := range bits[3:] {
				if strings.HasSuffix(child, ",") {
					child = child[0 : len(child)-1]
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

func (n *node) getValue() int {
	value := n.value
	for _, child := range n.children {
		value += child.getValue()
	}
	return value
}

func findUnbalanced(n *node) int {

	var values []int
	for _, child := range n.children {
		values = append(values, child.getValue())
	}

	if values[0] != values[1] && values[1] == values[2] {
		val := findUnbalanced(n.children[0])
		if val == -1 {
			return n.children[0].value - (values[0] - values[1])
		}
		return val
	}
	if values[0] != values[1] && values[0] == values[2] {
		val := findUnbalanced(n.children[1])
		if val == -1 {
			return n.children[1].value - (values[1] - values[0])
		}
		return val
	}
	if values[2] != values[1] && values[0] == values[1] {
		val := findUnbalanced(n.children[2])
		if val == -1 {
			return n.children[2].value - (values[2] - values[0])
		}
		return val
	}

	return -1
}

func getUnbalanced(data string) int {
	tree := buildTree(data)
	unbalanced := findUnbalanced(tree.root)
	return unbalanced
}

func (s *Server) Solve2017day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2017-7.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{StringAnswer: getBottom(trimmed)}, nil
}
