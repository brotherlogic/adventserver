package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"sort"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func makeKey(salt string, index int) string {
	fstring := fmt.Sprintf("%v%v", salt, index)
	hash := md5.Sum([]byte(fstring))
	return hex.EncodeToString(hash[:])
}

func countThrees(str string) string {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+1] && str[i] == str[i+2] {
			return string(str[i])
		}
	}

	return ""
}

func countFives(str string) []string {
	var res []string
	for i := 0; i < len(str)-4; i++ {
		if str[i] == str[i+1] && str[i] == str[i+2] && str[i] == str[i+3] && str[i] == str[i+4] {
			res = append(res, string(str[i]))
		}
	}

	return res
}

func buildKey(salt string, index int, fives map[string][]int) (string, map[string][]int) {
	key := makeKey(salt, index)
	threes := countThrees(key)

	if len(fives) == 0 {
		for i := index; i < index+1000; i++ {
			key := makeKey(salt, i)
			fivesList := countFives(key)
			for _, five := range fivesList {
				if _, ok := fives[five]; !ok {
					fives[five] = make([]int, 0)
				}
				log.Printf("%v -> %v", five, i)
				fives[five] = append(fives[five], i)
			}
		}
	} else {
		key := makeKey(salt, index+1000)
		fivesList := countFives(key)
		for _, five := range fivesList {
			if _, ok := fives[five]; !ok {
				fives[five] = make([]int, 0)
			}
			fives[five] = append(fives[five], index+1000)
		}
	}

	return threes, fives
}

func findFives(salt string) int {
	seen := make(map[int]string)
	index := 0
	count := 0
	var wins []int

	for {
		key := makeKey(salt, index)
		seen[index] = key
		c := countFives(key)
		for _, ff := range c {
			for i := index - 1001; i < index; i++ {
				if countThrees(seen[i]) == ff {
					wins = append(wins, i)
					count++
					if count > 64*2 {
						sort.Ints(wins)

						return wins[63]
					}
				}
			}
		}
		index++
	}
}

func buildKeys(salt string) map[int]int {
	res := make(map[int]int)
	fives := make(map[string][]int)

	index := 0
	curr := 1
	for len(res) < 64 {
		three, f := buildKey(salt, index, fives)
		fives = f

		found := false
		for _, fiv := range fives[three] {
			if fiv <= index+1+1000 && fiv > index {
				found = true
				break
			}
		}

		if found {
			res[curr] = index
			curr++
		}
		index++
	}

	return res
}

func (s *Server) Solve2016day14part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(findFives("ihaygndm"))}, nil
}
