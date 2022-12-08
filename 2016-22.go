package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type df struct {
	name                   string
	size, used, avail, use int
}

func convertNum(str string) int {
	num, _ := strconv.ParseInt(str[:len(str)-1], 10, 32)
	return int(num)
}

func buildNodes(data string) []*df {
	var nodes []*df
	for _, line := range strings.Split(data, "\n") {
		if line[0] == '/' {
			fields := strings.Fields(line)

			df := &df{name: fields[0]}
			df.size = convertNum(fields[1])
			df.used = convertNum(fields[1])
			df.avail = convertNum(fields[1])
			df.use = convertNum(fields[1])

			nodes = append(nodes, df)
		}
	}

	return nodes
}

func calcNodes(nodes []*df) int {
	count := 0
	for i, n1 := range nodes {
		for _, n2 := range nodes[i+1:] {
			if n1.used != 0 {
				if n1.used < n2.avail {
					count++
				}
			}
		}
	}

	return count
}

func (s *Server) Solve2016day22part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-22.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(calcNodes(buildNodes(data)))}, nil
}
