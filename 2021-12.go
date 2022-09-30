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

func countPaths(node *pnode, maxSeen int) int {
	count := 0
	for _, n := range node.neighbours {
		count += getPaths(n, make(map[string]int), maxSeen, "start-")
	}
	return count
}

func getPaths(node *pnode, seenSmall map[string]int, maxSeen int, sofar string) int {
	if node.name == "end" {
		//log.Printf("Seen %vend with %v", sofar, seenSmall)
		return 1
	}

	if node.name == "start" {
		return 0
	}

	if !node.bigcave {
		for seen, count := range seenSmall {
			if seen == node.name && count >= maxSeen {
				return 0
			}
		}

		seenSmall[node.name]++

		countBigs := 0
		for _, val := range seenSmall {
			if val > 1 {
				countBigs++
			}
		}
		if countBigs > 1 {
			return 0
		}
	}

	count := 0
	for _, child := range node.neighbours {
		count += getPaths(child, copyMap(seenSmall), maxSeen, sofar+node.name+"-")
	}

	return count
}

func copyMap(m map[string]int) map[string]int {
	nmap := make(map[string]int)
	for key, val := range m {
		nmap[key] = val
	}
	return nmap
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

	paths := countPaths(buildGraph(strings.TrimSpace(data)), 1)

	return &pb.SolveResponse{Answer: int32(paths)}, nil
}

func (s *Server) Solve2021day12part2(ctx context.Context) (*pb.SolveResponse, error) {
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

	paths := countPaths(buildGraph(strings.TrimSpace(data)), 2)

	return &pb.SolveResponse{Answer: int32(paths)}, nil
}
