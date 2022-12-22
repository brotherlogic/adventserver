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

func (p *program) resetCache() {
	for _, entry := range p.progs {
		entry.result = 0
	}
}

type entry struct {
	val         int64
	operator    string
	left, right string
	result      int64
	unknown     bool
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

func reverseResult(e *entry, value, result int64, isRight bool) int64 {
	switch e.operator {
	case "*":
		return result / value
	case "+":
		return result - value
	case "-":
		if isRight {
			return result + value
		} else {
			return value - result
		}
	default:
		if isRight {
			return result * value
		} else {
			return value / result
		}
	}
}

func findUnknown(prog *program, base string, result int64) int64 {

	if prog.progs[base].left == "" {
		return result
	}

	leftv, lun := evalProg(prog, prog.progs[base].left, "")
	rightv, run := evalProg(prog, prog.progs[base].right, "")

	if lun && run {
		log.Fatalf("UNSOLVABLE!")
	}

	if lun {
		ideal := reverseResult(prog.progs[base], rightv, result, true)
		return findUnknown(prog, prog.progs[base].left, ideal)
	} else {
		ideal := reverseResult(prog.progs[base], leftv, result, false)
		return findUnknown(prog, prog.progs[base].right, ideal)
	}
}

func evalProg(prog *program, ident, path string) (int64, bool) {
	node := prog.progs[ident]

	if node.unknown {
		return -1, true
	}

	if node.result > 0 {
		return node.result, false
	}
	if node.val > 0 {
		return node.val, false
	}

	switch node.operator {
	case "*":
		l, vl := evalProg(prog, node.left, path+"-"+node.left)
		r, vr := evalProg(prog, node.right, path+"-"+node.right)
		res := l * r
		if !vl && !vr {
			node.result = res
		}
		return res, vl || vr
	case "+":
		l, vl := evalProg(prog, node.left, path+"-"+node.left)
		r, vr := evalProg(prog, node.right, path+"-"+node.right)
		res := l + r
		if !vl && !vr {
			node.result = res
		}
		return res, vl || vr
	case "-":
		l, vl := evalProg(prog, node.left, path+"-"+node.left)
		r, vr := evalProg(prog, node.right, path+"-"+node.right)
		res := l - r
		if !vl && !vr {
			node.result = res
		}
		return res, vl || vr
	case "/":
		l, vl := evalProg(prog, node.left, path+"-"+node.left)
		r, vr := evalProg(prog, node.right, path+"-"+node.right)
		res := l / r
		if !vl && !vr {
			node.result = res
		}
		return res, vl || vr
	default:
		log.Fatalf("NOPE: %+v", node)
	}

	return -1, false
}

func (s *Server) Solve2022day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-21.txt")
	if err != nil {
		return nil, err
	}

	res, _ := evalProg(buildProgram(data), "root", "")
	return &pb.SolveResponse{BigAnswer: res}, nil
}

func (s *Server) Solve2022day21part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-21.txt")
	if err != nil {
		return nil, err
	}

	program := buildProgram(data)
	val, _ := evalProg(program, program.progs["root"].right, "")
	result := findUnknown(program, program.progs["root"].left, val)
	return &pb.SolveResponse{BigAnswer: result}, nil
}
