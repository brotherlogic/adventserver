package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type monkey struct {
	items       []int
	operation   string
	adjustment  int
	test        int
	trueMonkey  int
	falseMonkey int
	seen        int
}

func buildMonkeys(data string) []*monkey {
	return []*monkey{}
}

func runMonkeys(monkeys []*monkey) {

}

func getMonkeyTimes(monkeys []*monkey) []int {
	return []int{0, 0}
}

func (s *Server) Solve2022day11part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-11.txt")
	if err != nil {
		return nil, err
	}

	monkeys := buildMonkeys(data)
	for i := 0; i < 20; i++ {
		runMonkeys(monkeys)
	}
	vals := getMonkeyTimes(monkeys)

	return &pb.SolveResponse{Answer: int32(vals[0] * vals[1])}, nil
}
