package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func WorkRules(rules map[string]string, key string) uint {

	response := rules[key]

	if response == "" {
		response = key
	}

	if strings.Contains(response, "AND") {
		parts := strings.Split(response, " AND ")
		val1 := WorkRules(rules, parts[0])
		val2 := WorkRules(rules, parts[1])

		rules[parts[0]] = fmt.Sprintf("%v", val1)
		rules[parts[1]] = fmt.Sprintf("%v", val2)

		return val1 & val2
	}

	if strings.Contains(response, "OR") {
		parts := strings.Split(response, " OR ")
		val1 := WorkRules(rules, parts[0])
		val2 := WorkRules(rules, parts[1])

		rules[parts[0]] = fmt.Sprintf("%v", val1)
		rules[parts[1]] = fmt.Sprintf("%v", val2)

		return val1 | val2
	}

	if strings.Contains(response, "LSHIFT") {
		parts := strings.Split(response, " LSHIFT ")
		val1 := WorkRules(rules, parts[0])
		val2, _ := strconv.ParseUint(parts[1], 10, 8)

		rules[parts[0]] = fmt.Sprintf("%v", val1)

		return val1 << uint8(val2)
	}

	if strings.Contains(response, "RSHIFT") {
		parts := strings.Split(response, " RSHIFT ")
		val1 := WorkRules(rules, parts[0])
		val2, _ := strconv.ParseUint(parts[1], 10, 8)

		rules[parts[0]] = fmt.Sprintf("%v", val1)

		return val1 >> uint8(val2)
	}

	if strings.Contains(response, "NOT") {
		parts := strings.Split(response, "NOT ")
		val1 := uint16(WorkRules(rules, parts[1]))

		rules[parts[0]] = strconv.Itoa(int(val1))

		return uint(^val1)
	}

	m, err := regexp.MatchString(`\d+`, response)
	if err == nil && m {
		conv, _ := strconv.ParseInt(response, 10, 32)
		return uint(conv)
	}

	return WorkRules(rules, response)
}

func (s *Server) Solve2015day7part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-7.txt")
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(.*) -> (.*)`)
	var rules map[string]string
	var rules2 map[string]string
	rules = make(map[string]string)
	rules2 = make(map[string]string)

	for _, text := range strings.Split(strings.TrimSpace(data), "\n") {
		result := re.FindAllStringSubmatch(text, -1)
		rules[result[0][2]] = result[0][1]
		rules2[result[0][2]] = result[0][1]
	}

	return &pb.SolveResponse{Answer: int32(WorkRules(rules, "a"))}, nil
}

func (s *Server) Solve2015day7part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-7.txt")
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(.*) -> (.*)`)
	var rules map[string]string
	var rules2 map[string]string
	rules = make(map[string]string)
	rules2 = make(map[string]string)

	for _, text := range strings.Split(strings.TrimSpace(data), "\n") {
		result := re.FindAllStringSubmatch(text, -1)
		rules[result[0][2]] = result[0][1]
		rules2[result[0][2]] = result[0][1]
	}

	rules2["b"] = fmt.Sprintf("%v", WorkRules(rules, "a"))
	return &pb.SolveResponse{Answer: int32(WorkRules(rules2, "a"))}, nil
}
