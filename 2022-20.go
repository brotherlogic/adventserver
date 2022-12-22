package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type chain struct {
	value int64
	next  *chain
	prev  *chain
}

func printChain(head *chain) string {
	ret := fmt.Sprintf("%v", head.value)
	curr := head.next
	for curr != head {
		ret += fmt.Sprintf("|%v", curr.value)
		curr = curr.next
	}
	return ret
}

func unencrpyt(data string, mult int64, rep int) int64 {
	var runArr []*chain
	cHead := &chain{}
	curr := cHead
	var prev *chain

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			curr.value = int64(getInt32(line)) * mult
			runArr = append(runArr, curr)
			if prev != nil {
				curr.prev = prev
				prev.next = curr
			}
			prev = curr
			curr = &chain{}
		}
	}
	cHead.prev = prev
	cHead.prev.next = cHead

	curr = cHead
	for vv := 0; vv < rep; vv++ {
		for _, v := range runArr {
			if v.value != 0 {
				curr = v
				val := curr.value
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
				add := curr
				if curr.value > 0 {
					for i := int64(0); i < val%int64(len(runArr)-1); i++ {
						add = add.next
					}
					curr.next = add.next
					curr.next.prev = curr
					add.next = curr
					curr.prev = add

				} else if val < 0 {
					for i := int64(0); i < -val%int64(len(runArr)-1); i++ {
						add = add.prev
					}
					curr.prev = add.prev
					curr.prev.next = curr
					add.prev = curr
					curr.next = add
				}
			}
		}
	}

	var zero *chain
	for {
		if curr.value == 0 {
			zero = curr
			break
		}
		curr = curr.next
	}

	value := int64(0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			zero = zero.next
		}

		value += zero.value
	}

	return value
}

func printNumMap(numMap map[int]int) string {
	narr := make([]int, len(numMap))
	for key, val := range numMap {
		narr[val] = key
	}
	return fmt.Sprintf("%v", narr)
}

func moveNumber(arr []int, num int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = i
	}

	numMap = moveMap(numMap, num)

	newarr := make([]int, len(arr))
	for key, val := range numMap {
		newarr[val] = key
	}

	return newarr
}

func moveMap(numMap map[int]int, num int) map[int]int {
	if num == 0 {
		return numMap
	}

	nIndex := numMap[num]
	llen := len(numMap)

	newIndex := (nIndex + num) % llen
	if num < 0 {
		newIndex = (nIndex + num - 1) % llen
	}
	if newIndex < 0 {
		newIndex = llen + newIndex
	}

	following := 0
	for key, val := range numMap {
		if val == newIndex {
			following = key
		}
	}

	// Remove the number
	delete(numMap, num)
	for key, val := range numMap {
		if val > nIndex {
			numMap[key] = val - 1
		}
	}

	newNewIndex := numMap[following] + 1

	for key, val := range numMap {
		if val >= newNewIndex {
			numMap[key] = val + 1
		}
	}
	numMap[num] = newNewIndex

	return numMap
}

func (s *Server) Solve2022day20part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-20.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(unencrpyt(data, 1, 1))}, nil
}

func (s *Server) Solve2022day20part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-20.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{BigAnswer: (unencrpyt(data, 811589153, 10))}, nil
}
