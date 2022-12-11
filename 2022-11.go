package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type monkey struct {
	number      int
	items       []int
	operation   string
	adjustment  int
	test        int
	trueMonkey  int
	falseMonkey int
	seen        int
}

func buildMonkeys(data string) []*monkey {
	var cmonkey *monkey
	var monkeys []*monkey
	for _, line := range strings.Split(data, "\n") {
		nline := strings.TrimSpace(line)
		if len(nline) > 0 {
			fields := strings.Fields(nline)
			switch fields[0] {
			case "Monkey":
				num, _ := strconv.ParseInt(fields[1][:len(fields[1])-2], 10, 32)
				if num > 0 {
					monkeys = append(monkeys, cmonkey)
				}
				cmonkey = &monkey{number: int(num), items: make([]int, 0)}

			case "Starting":
				tline := strings.Split(nline, ":")
				nums := strings.Split(tline[1], ",")
				for _, num := range nums {
					numv, _ := strconv.ParseInt(num, 10, 32)
					cmonkey.items = append(cmonkey.items, int(numv))
				}
			case "Operation":
				cmonkey.operation = fields[4]
				num, _ := strconv.ParseInt(fields[5], 10, 32)
				cmonkey.adjustment = int(num)
			case "Test":
				num, _ := strconv.ParseInt(fields[3], 10, 32)
				cmonkey.test = int(num)
			case "If":
				num, _ := strconv.ParseInt(fields[5], 10, 32)
				if fields[1] == "true:" {
					cmonkey.trueMonkey = int(num)
				} else {
					cmonkey.falseMonkey = int(num)
				}
			}
		}
	}

	monkeys = append(monkeys, cmonkey)

	return monkeys
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
