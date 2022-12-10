package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func ropeMove(hx, hy, tx, ty int) (int, int) {
	if hx == tx {
		if hy-ty > 1 {
			return tx, ty + 1
		}
		if ty-hy > 1 {
			return tx, ty - 1
		}

		return tx, ty
	}

	if hy == ty {
		if hx-tx > 1 {
			return tx + 1, ty
		}

		if tx-hx > 1 {
			return tx - 1, ty
		}
	}

	if (hx-tx == 2 && ty-hy == 1) ||
		(ty-hy == 2 && hx-tx == 1) {
		return tx + 1, ty - 1
	}

	if (hx-tx == 2 && hy-ty == 1) ||
		(hy-ty == 2 && hx-tx == 1) {
		return tx + 1, ty + 1
	}

	if (tx-hx == 2 && ty-hy == 1) ||
		(ty-hy == 2 && tx-hx == 1) {
		return tx - 1, ty - 1
	}

	if (tx-hx == 2 && hy-ty == 1) ||
		(hy-ty == 2 && tx-hx == 1) {
		return tx - 1, ty + 1
	}

	return tx, ty
}

func runRopeBridge(data string, num int) int {
	friends := make([][]int, num)
	for i := range friends {
		friends[i] = make([]int, 2)
	}
	seen := make(map[int]map[int]bool)
	seen[0] = make(map[int]bool)
	seen[0][0] = true

	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			fields := strings.Fields(line)
			count, _ := strconv.ParseInt(fields[1], 10, 32)

			xadj, yadj := 0, 0
			switch fields[0] {
			case "U":
				yadj = -1
			case "D":
				yadj = 1
			case "L":
				xadj = -1
			case "R":
				xadj = 1
			}

			for i := 0; i < int(count); i++ {
				friends[0][0] += xadj
				friends[0][1] += yadj

				for i := 1; i < num; i++ {
					nx, ny := ropeMove(friends[i-1][0], friends[i-1][1], friends[i][0], friends[i][1])
					friends[i][0] = nx
					friends[i][1] = ny
				}

				if _, ok := seen[friends[len(friends)-1][0]]; !ok {
					seen[friends[len(friends)-1][0]] = make(map[int]bool)
				}
			}
			seen[friends[len(friends)-1][0]][friends[len(friends)-1][1]] = true
		}
	}

	cc := 0
	for _, blah := range seen {
		for _, blahblah := range blah {
			if blahblah {
				cc++
			}
		}
	}

	return cc
}

func (s *Server) Solve2022day9part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-9.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runRopeBridge(data, 2))}, nil
}
