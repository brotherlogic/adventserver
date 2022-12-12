package main

import (
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type monkey struct {
	number      int
	items       []int64
	operation   string
	adjustment  int64
	test        int64
	trueMonkey  int
	falseMonkey int
	seen        int64
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
				num, _ := strconv.ParseInt(fields[1][:len(fields[1])-1], 10, 64)
				if num > 0 {
					monkeys = append(monkeys, cmonkey)
				}
				cmonkey = &monkey{number: int(num), items: make([]int64, 0)}

			case "Starting":
				tline := strings.Split(nline, ":")
				if len(strings.TrimSpace(tline[1])) > 0 {
					nums := strings.Split(tline[1], ",")
					for _, num := range nums {
						numv, _ := strconv.ParseInt(strings.TrimSpace(num), 10, 64)
						cmonkey.items = append(cmonkey.items, int64(numv))
					}
				}
			case "Operation:":
				cmonkey.operation = fields[4]
				num, _ := strconv.ParseInt(fields[5], 10, 64)
				if num == 0 {
					cmonkey.adjustment = -1
				} else {
					cmonkey.adjustment = int64(num)
				}
			case "Test:":
				num, _ := strconv.ParseInt(fields[3], 10, 64)
				cmonkey.test = int64(num)
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

func runMonkeysLong(monkeys []*monkey) {
	mmap := make(map[int]*monkey)
	controller := int64(1)
	for _, monkey := range monkeys {
		mmap[monkey.number] = monkey
		controller *= monkey.test
	}

	for _, monkey := range monkeys {
		for len(monkey.items) > 0 {
			vitem := monkey.items[0]

			vitem = vitem % controller
			if vitem == 0 {
				vitem = monkey.test
			}

			switch monkey.operation {
			case "*":
				if monkey.adjustment > 0 {
					vitem *= monkey.adjustment
				} else {
					vitem = vitem * vitem
				}
			case "+":
				if monkey.adjustment > 0 {
					vitem += monkey.adjustment
				} else {
					vitem = vitem + vitem
				}
			}

			if vitem%monkey.test == 0 {
				mmap[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, vitem)
			} else {
				mmap[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, vitem)
			}
			monkey.items = monkey.items[1:]
			monkey.seen++
		}

	}
}

func runMonkeys(monkeys []*monkey) {
	mmap := make(map[int]*monkey)
	for _, monkey := range monkeys {
		mmap[monkey.number] = monkey
	}

	for _, monkey := range monkeys {
		for len(monkey.items) > 0 {
			vitem := monkey.items[0]
			switch monkey.operation {
			case "*":
				if monkey.adjustment > 0 {
					vitem *= monkey.adjustment
				} else {
					vitem = vitem * vitem
				}
			case "+":
				if monkey.adjustment > 0 {
					vitem += monkey.adjustment
				} else {
					vitem = vitem + vitem
				}
			}

			vitem = vitem / 3

			if vitem%monkey.test == 0 {
				mmap[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, vitem)
			} else {
				mmap[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, vitem)
			}
			monkey.items = monkey.items[1:]
			monkey.seen++
		}

	}
}

func getMonkeyTimes(monkeys []*monkey) []int64 {

	var values []int64
	for _, m := range monkeys {
		values = append(values, m.seen)
	}

	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })

	return values
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

func (s *Server) Solve2022day11part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-11.txt")
	if err != nil {
		return nil, err
	}

	monkeys := buildMonkeys(data)
	for i := 0; i < 10000; i++ {
		runMonkeysLong(monkeys)
	}
	vals := getMonkeyTimes(monkeys)

	return &pb.SolveResponse{Answer: int32(vals[0] * vals[1])}, nil
}
