package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type mazeEntry struct {
	x, y int
	key  string
	path string
}

func rallowed(r byte) bool {
	return r == 'b' ||
		r == 'c' ||
		r == 'd' ||
		r == 'e' ||
		r == 'f'
}

func (m mazeEntry) nextSteps() []mazeEntry {
	hash := getHash(m.key + m.path)

	var follow []mazeEntry
	if rallowed(hash[0]) && m.y > 0 {
		follow = append(follow, mazeEntry{m.x, m.y - 1, m.key, m.path + "U"})
	}
	if rallowed(hash[1]) && m.y < 3 {
		follow = append(follow, mazeEntry{m.x, m.y + 1, m.key, m.path + "D"})
	}
	if rallowed(hash[2]) && m.x > 0 {
		follow = append(follow, mazeEntry{m.x - 1, m.y, m.key, m.path + "L"})
	}
	if rallowed(hash[3]) && m.x < 3 {
		follow = append(follow, mazeEntry{m.x + 1, m.y, m.key, m.path + "R"})
	}

	return follow

}

func getShortestPath(key string) string {
	searchQueue := []mazeEntry{{0, 0, key, ""}}

	for {
		head := searchQueue[0]
		searchQueue = searchQueue[1:]

		if head.x == 3 && head.y == 3 {
			return head.path
		}

		searchQueue = append(searchQueue, head.nextSteps()...)
	}
}

func (s *Server) Solve2016day17part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{StringAnswer: (getShortestPath("mmsxrhfx"))}, nil
}
