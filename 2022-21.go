package main

import (
	"log"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type program struct {
	progs map[string]*entry
}

type entry struct {
	val         int64
	operator    string
	left, right string
	result      int64
}

func buildLine(line string) (string, *entry) {
	fields := strings.Fields(strings.ReplaceAll(line, ":", ""))
	if len(fields) == 2 {
		value := getInt32(fields[1])
		return fields[0], &entry{val: int64(value)}
	}

	return fields[0], &entry{left: fields[1], operator: fields[2], right: fields[3]}
}

func buildProgram(data string) *program {
	prog := &program{progs: make(map[string]*entry)}
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		name, ent := buildLine(strings.TrimSpace(line))
		prog.progs[name] = ent
	}
	return prog
}

func evalProg(prog *program, ident string) int64 {
	node := prog.progs[ident]
	if node.result > 0 {
		return node.result
	}
	if node.val > 0 {
		return node.val
	}

	switch node.operator {
	case "*":
		res := evalProg(prog, node.left) * evalProg(prog, node.right)
		node.result = res
		return res
	case "+":
		res := evalProg(prog, node.left) + evalProg(prog, node.right)
		node.result = res
		return res
	case "-":
		res := evalProg(prog, node.left) - evalProg(prog, node.right)
		node.result = res
		return res
	case "/":
		res := evalProg(prog, node.left) / evalProg(prog, node.right)
		node.result = res
		return res
	default:
		log.Fatalf("NOPE: %+v", node)
	}

	return -1
}

func (s *Server) Solve2022day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-21.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{BigAnswer: (evalProg(buildProgram(data), "root"))}, nil
}
