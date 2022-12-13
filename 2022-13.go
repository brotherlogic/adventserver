package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type lelem struct {
	value int
	elems []*lelem
}

func printLelem(lelem *lelem) string {
	if lelem.value >= 0 {
		return fmt.Sprintf("%v", lelem.value)
	}

	ret := "["
	for _, val := range lelem.elems {
		ret += fmt.Sprintf("%v,", printLelem(val))
	}

	if len(lelem.elems) == 0 {
		return ret + "]"
	}

	return ret[:len(ret)-1] + "]"
}

func buildLelem(pointer int, line string) (*lelem, int) {
	ret := &lelem{elems: make([]*lelem, 0), value: -1}

	for pointer < len(line) {
		if line[pointer] == '[' {
			nelem, npoint := buildLelem(pointer+1, line)
			ret.elems = append(ret.elems, nelem)
			pointer = npoint
		} else if line[pointer] == ']' {
			return ret, pointer + 1
		} else if line[pointer] == ',' {
			pointer++
		} else {
			val, _ := strconv.ParseInt(string(line[pointer]), 10, 32)
			if line[pointer] == '1' && line[pointer+1] == '0' {
				val = int64(10)
			}

			ret.elems = append(ret.elems, &lelem{value: int(val)})
			pointer++
		}
	}

	return ret.elems[0], pointer
}

func rightOrder(l1, l2 *lelem) int {
	if l1.value >= 0 && l2.value >= 0 {
		if l1.value < l2.value {
			return 1
		}
		if l1.value > l2.value {
			return -1
		}

		return 0
	}

	if l1.value >= 0 {
		return rightOrder(&lelem{value: -1, elems: []*lelem{&lelem{value: l1.value}}}, l2)
	}

	if l2.value >= 0 {
		return rightOrder(l1, &lelem{value: -1, elems: []*lelem{&lelem{value: l2.value}}})
	}

	for i := 0; i < len(l1.elems); i++ {
		if i >= len(l2.elems) {
			return -1
		}

		comp := rightOrder(l1.elems[i], l2.elems[i])
		if comp != 0 {
			return comp
		}
	}

	return 1
}

func computeIndexSum(data string) int {
	elems := strings.Split(data, "\n")

	sumv := 0
	for i := 0; i < len(elems); i += 3 {
		l1, _ := buildLelem(0, strings.TrimSpace(elems[i]))
		l2, _ := buildLelem(0, strings.TrimSpace(elems[i+1]))
		log.Printf("%v vs %v => %v", printLelem(l1), printLelem(l2), rightOrder(l1, l2))

		if rightOrder(l1, l2) == 1 {
			sumv += i/3 + 1
		}
	}

	return sumv
}

func (s *Server) Solve2022day13part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-13.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(computeIndexSum(data))}, nil
}
