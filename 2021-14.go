package main

import (
	"math"
	"math/big"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildRules(data string) map[string]string {
	rules := make(map[string]string)

	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "->") {
			elems := strings.Split(line, "->")
			rules[strings.TrimSpace(elems[0])] = strings.TrimSpace(elems[1])
		}
	}
	return rules
}

func convertToMap(line string) map[string]*big.Int {
	mapper := make(map[string]*big.Int)
	for i := 0; i < len(line)-1; i++ {
		if val, ok := mapper[string(line[i])+string(line[i+1])]; ok {
			val.Add(val, big.NewInt(1))
		} else {
			mapper[string(line[i])+string(line[i+1])] = big.NewInt(1)
		}
	}
	return mapper
}

func runData(data string, count int) map[string]*big.Int {
	first := strings.Split(data, "\n")[0]
	mapper := convertToMap(strings.TrimSpace(first))
	rules := buildRules(data)

	for i := 0; i < count; i++ {
		mapper = runRules(mapper, rules)
	}

	return mapper
}

func runRules(mapper map[string]*big.Int, rules map[string]string) map[string]*big.Int {
	nmap := make(map[string]*big.Int)
	for val, count := range mapper {
		if news, ok := rules[val]; ok {
			if v, ok := nmap[string(val[0])+news]; ok {
				v.Add(v, count)
			} else {
				nmap[string(val[0])+news] = new(big.Int)
				nmap[string(val[0])+news].Set(count)
			}
			if v, ok := nmap[news+string(val[1])]; ok {
				v.Add(v, count)
			} else {
				nmap[news+string(val[1])] = new(big.Int)
				nmap[news+string(val[1])].Set(count)
			}
		}
	}
	return nmap
}

func getCommons(line map[string]*big.Int) (*big.Int, *big.Int) {
	counts := make(map[string]*big.Int)
	for str, co := range line {
		if val, ok := counts[string(str[0])]; ok {
			val.Add(val, co)
		} else {
			counts[string(str[0])] = new(big.Int)
			counts[string(str[0])].Set(co)
		}
		if val, ok := counts[string(str[1])]; ok {
			val.Add(val, co)
		} else {
			counts[string(str[1])] = new(big.Int)
			counts[string(str[1])].Set(co)
		}
	}

	highest := big.NewInt(0)
	lowest := big.NewInt(math.MaxInt64)

	for _, v := range counts {
		if v.Cmp(highest) > 0 {
			highest = v
			//log.Printf("HIGH %v -> %v", key, (v+1)/2)
		}
		if v.Cmp(lowest) < 0 {
			lowest = v
		}
	}

	return (highest.Add(highest, big.NewInt(1)).Div(highest, big.NewInt(2))), (lowest.Add(lowest, big.NewInt(1)).Div(lowest, big.NewInt(2)))
}

func (s *Server) Solve2021day14part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-14.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	newone := runData(trimmed, 10)
	mc, lc := getCommons(newone)
	return &pb.SolveResponse{StringAnswer: mc.Sub(mc, lc).String()}, nil
}

func (s *Server) Solve2021day14part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-14.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	newone := runData(trimmed, 40)
	mc, lc := getCommons(newone)
	return &pb.SolveResponse{StringAnswer: mc.Sub(mc, lc).String()}, nil
}
