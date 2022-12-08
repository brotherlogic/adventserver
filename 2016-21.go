package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func replace(s1, s2, code string) string {
	return strings.Replace(strings.Replace(strings.Replace(code, s1, "1", 1), s2, s1, 1), "1", s2, 1)
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func translateCode(code, trans string) string {
	fields := strings.Fields(trans)
	switch fields[0] {
	case "swap":
		if fields[1] == "position" {
			p1, _ := strconv.ParseInt(fields[2], 10, 32)
			p2, _ := strconv.ParseInt(fields[5], 10, 32)

			return replace(string(code[int(p1)]), string(code[int(p2)]), code)
		}
		if fields[1] == "letter" {
			return replace(fields[2], fields[5], code)
		}
	case "rotate":
		if fields[1] == "left" {
			count, _ := strconv.ParseInt(fields[2], 10, 32)
			return code[count:] + code[:count]
		}

		if fields[1] == "right" {
			count, _ := strconv.ParseInt(fields[2], 10, 32)
			nv := len(code) - int(count)
			return code[nv:] + code[:nv]
		}

		if fields[1] == "based" {
			index := strings.Index(code, string(fields[6]))
			if index >= 4 {
				index++
			}
			ncode := code[len(code)-1:] + code[:len(code)-1]
			return ncode[len(code)-index:] + ncode[:len(code)-index]
		}
	case "reverse":
		p1, _ := strconv.ParseInt(fields[2], 10, 32)
		p2, _ := strconv.ParseInt(fields[4], 10, 32)

		return code[:p1] + reverse(code[p1:p2+1]) + code[p2+1:]
	case "move":
		p1, _ := strconv.ParseInt(fields[2], 10, 32)
		p2, _ := strconv.ParseInt(fields[5], 10, 32)

		val := string(code[p1])
		nstr := code[:p1] + code[p1+1:]
		return nstr[:p2] + val + nstr[p2:]

	}
	return code
}

func fullTranslate(data, code string) string {
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			code = translateCode(code, strings.TrimSpace(line))
		}
	}
	return code
}

func generateAll(code, sofar, goal, data string) string {
	if len(code) == 0 {
		ncode := fullTranslate(data, sofar)
		if ncode == goal {
			return sofar
		}
		return ""
	}

	for _, char := range code {
		val := generateAll(strings.ReplaceAll(code, string(char), ""), sofar+string(char), goal, data)
		if val != "" {
			return val
		}
	}

	return ""
}

func (s *Server) Solve2016day21part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-21.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: fullTranslate(data, "abcdefgh")}, nil
}

func (s *Server) Solve2016day21part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2016-21.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: fullTranslate(data, "abcdefgh")}, nil
}
