package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type file struct {
	name string
	size int
}

type dir struct {
	name  string
	dirs  []*dir
	files []*file
}

func buildDirs(data string) *dir {
	return &dir{name: "/", dirs: make([]*dir, 0), files: make([]*file, 0)}
}

func (d *dir) dirSum(limit int) int {
	return 0
}

func (s *Server) Solve2022day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-7.txt")
	if err != nil {
		return nil, err
	}

	d := buildDirs(data)

	return &pb.SolveResponse{Answer: int32(d.dirSum(100000))}, nil
}
