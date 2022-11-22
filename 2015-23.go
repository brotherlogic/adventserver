package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type computer struct {
	a, b int
}

func runProgram(program string) computer {
	commands := strings.Split(program, "\n")
	for i := range commands {
		commands[i] = strings.TrimSpace(commands[i])
	}

	curr := computer{}
	i := 0
	for i < len(commands) {
		elems := strings.Fields(commands[i])
		if len(elems) == 0 {
			return curr
		}
		switch elems[0] {
		case "hlf":
			if elems[1] == "a" {
				curr.a = curr.a / 2
			} else {
				curr.b = curr.b / 2
			}
			i++
		case "tpl":
			if elems[1] == "a" {
				curr.a = curr.a * 3
			} else {
				curr.b = curr.b * 3
			}
			i++
		case "inc":
			if elems[1] == "a" {
				curr.a = curr.a + 1
			} else {
				curr.b = curr.b + 1
			}
			i++
		case "jmp":
			val, _ := strconv.ParseInt(elems[1], 10, 16)
			i += int(val)
		case "jio":
			if elems[1] == "a," && curr.a%2 == 1 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else if elems[1] == "b," && curr.b%2 == 1 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else {
				i++
			}
		case "jie":
			if elems[1] == "a," && curr.a%2 == 0 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else if elems[1] == "b," && curr.b%2 == 0 {
				val, _ := strconv.ParseInt(elems[2], 10, 16)
				i += int(val)
			} else {
				i++
			}
		}
	}

	return curr
}

func (s *Server) Solve2015day23part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-23.txt")
	if err != nil {
		return nil, err
	}

	result := runProgram(data)
	return &pb.SolveResponse{Answer: int32(result.b)}, nil
}
