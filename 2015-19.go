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

func buildString(parts []string, pos int, from, adj string) string {
	nstr := ""
	for i := 0; i < pos; i++ {
		nstr += parts[i] + from
	}
	nstr += parts[pos] + adj
	for i := pos + 1; i < len(parts)-1; i++ {
		nstr += parts[i] + from
	}
	nstr += parts[len(parts)-1]
	return nstr
}

func translate(key string, trans map[string][]string) map[string]bool {
	result := make(map[string]bool)

	for adj, tos := range trans {
		for _, to := range tos {
			parts := strings.Split(key, adj)
			for i := 0; i < len(parts)-1; i++ {
				nstr := buildString(parts, i, adj, to)
				result[nstr] = true
			}
		}
	}

	//log.Printf("RES %v", result)
	return result
}

func treeMolecules(data string) int {
	return 0
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
