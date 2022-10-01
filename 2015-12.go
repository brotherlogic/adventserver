package main

import (
	"encoding/json"
	"log"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func countJson(s interface{}) int {
	switch v := s.(type) {
	case float64:
		return int(v)
	case []interface{}:
		return countArr(v)
	case map[string]interface{}:
		return countObj(v)
	default:
		log.Fatalf("Don't know what to do with %T", v)
	}

	return 0
}

func countArrStr(s string) (int, error) {
	var result []interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return 0, err
	}
	return countArr(result), nil
}

func countArr(result []interface{}) int {
	sumv := 0
	for _, value := range result {
		sumv += countJson(value)
	}

	return sumv
}

func countObjStr(s string) (int, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return 0, err
	}
	return countObj(result), nil
}

func countObj(result map[string]interface{}) int {
	sumv := 0
	for _, value := range result {
		sumv += countJson(value)
	}

	return sumv
}

func countNumbers(s string) int {
	sv, err := countArrStr(s)
	if err == nil {
		return sv
	}

	sc, err := countObjStr(s)
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
		ans += int32(countNumbers(str))
	}

	return &pb.SolveResponse{Answer: ans}, nil
}
