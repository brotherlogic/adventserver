package main

import (
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildMaps(data string) (map[string][]string, string) {
	done := false
	mapper := make(map[string][]string)
	key := ""
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			done = true
		} else {
			if !done {
				elems := strings.Fields(line)
				mapper[elems[0]] = append(mapper[elems[0]], elems[2])
			} else {
				key = line
			}
		}
	}

	return mapper, key
}

func translate(key string, trans map[string][]string) map[string]bool {
	result := make(map[string]bool)

	for i, ch := range key {
		if res, ok := trans[string(ch)]; ok {
			for _, r := range res {
				nstr := key[:i] + r + key[i:]
				result[nstr] = true
			}
		}
	}

	return result
}

func getMolecules(data string) int {
	trans, key := buildMaps(data)

	res := translate(key, trans)

	return len(res)
}

func (s *Server) Solve2015day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-19.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(getMolecules(trimmed))}, nil
}
