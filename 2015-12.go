package main

import (
	"encoding/json"
	"log"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countJson(s interface{}, skip string) (int, bool) {
	switch v := s.(type) {
	case float64:
		return int(v), false
	case []interface{}:
		return countArr(v, skip), false
	case map[string]interface{}:
		return countObj(v, skip), false
	case string:
		return 0, v == skip
	default:
		log.Fatalf("Don't know what to do with %T", v)
	}

	return 0, false
}

func countArrStr(s string, skip string) (int, error) {
	var result []interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return 0, err
	}
	return countArr(result, skip), nil
}

func countArr(result []interface{}, skip string) int {
	sumv := 0
	for _, value := range result {
		sum, _ := countJson(value, skip)
		sumv += sum
	}

	return sumv
}

func countObjStr(s string, skip string) (int, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return 0, err
	}
	return countObj(result, skip), nil
}

func countObj(result map[string]interface{}, skip string) int {
	sumv := 0
	for _, value := range result {
		sum, skip := countJson(value, skip)
		if skip {
			return 0
		}
		sumv += sum
	}

	return sumv
}

func countNumbers(s string, skip string) int {
	sv, err := countArrStr(s, skip)
	if err == nil {
		return sv
	}

	sc, err := countObjStr(s, skip)
	if err == nil {
		return sc
	}

	return 0
}

func (s *Server) Solve2015day12part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-12.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	ans := int32(0)
	for _, str := range strings.Split(trimmed, "\n") {
		ans += int32(countNumbers(str, "NOTHING"))
	}

	return &pb.SolveResponse{Answer: ans}, nil
}

func (s *Server) Solve2015day12part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-12.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	ans := int32(0)
	for _, str := range strings.Split(trimmed, "\n") {
		ans += int32(countNumbers(str, "red"))
	}

	return &pb.SolveResponse{Answer: ans}, nil
}
