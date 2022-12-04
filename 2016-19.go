package main

import (
	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func runPresents(num int) int {
	elfNum := make([]int, num)
	elfPres := make([]int, num)
	for i := 0; i < num; i++ {
		elfNum[i] = i + 1
		elfPres[i] = 1
	}

	pointer := 0
	for {
		if len(elfNum) == 1 {
			return elfNum[0]
		}

		if elfPres[pointer] == 0 {
			var nElfNum []int
			var nElfPres []int
			for i := 0; i < len(elfNum); i++ {
				if i != pointer {
					nElfNum = append(nElfNum, elfNum[i])
					nElfPres = append(nElfPres, elfPres[i])
				}
			}

			elfNum = nElfNum
			elfPres = nElfPres
		} else {
			elfPres[pointer] += elfPres[(pointer+1)%len(elfPres)]
			elfPres[(pointer+1)%len(elfPres)] = 0
			pointer++
			pointer = pointer % len(elfPres)
		}
	}

}

func (s *Server) Solve2016day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(runPresents(3018458))}, nil
}
