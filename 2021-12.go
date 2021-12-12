package main

import (
	"log"
	"strings"
	"unicode"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type pnode struct {
	name       string
	neighbours []*pnode
	bigcave    bool
}

func buildGraph(data string) *pnode {
	var nodes []*pnode

	for _, line := range strings.Split(data, "\n") {
		elems := strings.Split(strings.TrimSpace(line), "-")

		var stn, enn *pnode
		for _, n := range nodes {
			if n.name == elems[0] {
				stn = n
			}
			if n.name == elems[1] {
				enn = n
			}
		}

		if stn == nil {
			stn = &pnode{name: elems[0]}
			nodes = append(nodes, stn)
		}
		if enn == nil {
			enn = &pnode{name: elems[1]}
			nodes = append(nodes, enn)
		}

		stn.neighbours = append(stn.neighbours, enn)
		enn.neighbours = append(enn.neighbours, stn)
	}

	for _, node := range nodes {
		node.bigcave = unicode.IsUpper(rune(node.name[0]))
	}

	for _, node := range nodes {
		if node.name == "start" {
			return node
		}
	}

	log.Printf("START NOT FOUND")
	return nil
}

func countPaths(node *pnode) int {
	count := 0
	for _, n := range node.neighbours {
		count += getPaths(n, make([]string, 0), "")
	}
	return count
}

func getPaths(node *pnode, seenSmall []string, sofar string) int {
	if node.name == "end" {
		return 1
	}

	if node.name == "start" {
		return 0
	}

	if !node.bigcave {
		for _, seen := range seenSmall {
			if seen == node.name {
				return 0
			}
		}

		seenSmall = append(seenSmall, node.name)
	}

	count := 0
	for _, child := range node.neighbours {
		count += getPaths(child, seenSmall, sofar+node.name)
	}

	return count
}

func (s *Server) Solve2021day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data := `start-YA
	ps-yq
	zt-mu
	JS-yi
	yq-VJ
	QT-ps
	start-yq
	YA-yi
	start-nf
	nf-YA
	nf-JS
	JS-ez
	yq-JS
	ps-JS
	ps-yi
	yq-nf
	QT-yi
	end-QT
	nf-yi
	zt-QT
	end-ez
	yq-YA
	end-JS`

	paths := countPaths(buildGraph(strings.TrimSpace(data)))

	return &pb.SolveResponse{Answer: int32(paths)}, nil
}
