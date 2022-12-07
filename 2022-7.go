package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type file struct {
	name string
	size int
}

type dir struct {
	name   string
	dirs   []*dir
	files  []*file
	parent *dir
}

func buildDirs(data string) *dir {
	head := &dir{name: "/", dirs: make([]*dir, 0), files: make([]*file, 0), parent: nil}

	curr := head
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		fields := strings.Fields(line)
		switch fields[0] {
		case "$":
			if fields[1] == "cd" {
				if fields[2] == "/" {
					curr = head
					continue
				}

				if fields[2] == ".." {
					curr = curr.parent
					continue
				}

				found := false
				for _, dd := range curr.dirs {
					if dd.name == fields[2] {
						curr = dd
						found = true
					}
				}
				if found {
					continue
				}

				log.Fatalf("Unable to locate directory: %v", line)
			}

		case "dir":
			curr.dirs = append(curr.dirs, &dir{name: fields[1], dirs: make([]*dir, 0), files: make([]*file, 0), parent: curr})

		default:
			size, _ := strconv.ParseInt(fields[0], 10, 32)
			curr.files = append(curr.files, &file{name: fields[1], size: int(size)})
		}
	}

	return head
}

func (d *dir) dirSum(limit int) int {
	total := 0

	for _, dd := range d.dirs {
		if dd.size() <= limit {
			total += dd.size()
		}

		total += dd.dirSum(limit)
	}

	return total
}

func (d *dir) size() int {

	local := 0
	for _, f := range d.files {
		local += f.size
	}

	for _, dd := range d.dirs {
		local += dd.size()
	}

	return local
}

func (d *dir) remove(total int, left int) int {

	used := d.size()
	actLeft := total - used
	toRemove := left - actLeft

	smallest := math.MaxInt
	for _, dd := range d.dirs {
		smol := dd.small(toRemove)
		if smol < smallest {
			smallest = smol
		}
	}

	return smallest
}

func (d *dir) small(toRemove int) int {
	smallest := math.MaxInt

	smaller := d.size()
	if smaller < smallest && smaller >= toRemove {
		smallest = smaller
	}

	for _, dd := range d.dirs {
		sm := dd.small(toRemove)
		if sm < smallest {
			smallest = sm
		}
	}

	return smallest

}

func (s *Server) Solve2022day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-7.txt")
	if err != nil {
		return nil, err
	}

	d := buildDirs(data)

	return &pb.SolveResponse{Answer: int32(d.dirSum(100000))}, nil
}

func (s *Server) Solve2022day7part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-7.txt")
	if err != nil {
		return nil, err
	}

	d := buildDirs(data)

	return &pb.SolveResponse{Answer: int32(d.remove(70000000, 30000000))}, nil
}
