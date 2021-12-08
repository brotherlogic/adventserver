package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func in(val int, arr []int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}

	return false
}

func reduceTo(arr []string, word string) []string {
	var newa []string
	for _, char := range arr {
		if strings.Contains(word, string(char)) {
			newa = append(newa, string(char))
		}
	}

	return newa
}

func remove(arr []string, word string) []string {
	var newa []string
	for _, char := range arr {
		if !strings.Contains(word, string(char)) {
			newa = append(newa, string(char))
		}
	}

	return newa
}

func reduce(theory map[int][]string, ind []int, word string) {
	for key, arr := range theory {
		if in(key, ind) {
			theory[key] = reduceTo(arr, word)
		} else {
			theory[key] = remove(arr, word)
		}
	}
}

func buildOverlap(arr []string) string {
	counts := make(map[string]int)
	for _, word := range arr {
		for _, char := range word {
			counts[string(char)]++
		}
	}

	word := ""
	for char, count := range counts {
		if count == len(arr) {
			word += char
		}
	}

	return word
}

func invert(theory map[int][]string) map[string]int {
	mapper := make(map[string]int)
	for key, vals := range theory {
		mapper[vals[0]] = key
	}
	return mapper
}

func resolve(theory map[string]int, words string) int {
	val := ""
	for _, word := range strings.Fields(words) {
		var blah []int
		for _, c := range word {
			blah = append(blah, theory[string(c)])
		}

		sort.Ints(blah)
		num := -1
		switch len(blah) {
		case 2:
			num = 1
		case 3:
			num = 7
		case 4:
			num = 4
		case 5:
			if blah[1] == 1 {
				num = 5
			} else if blah[3] == 4 {
				num = 2
			} else {
				num = 3
			}
		case 6:
			if blah[2] == 3 {
				num = 6
			} else if blah[3] == 4 {
				num = 0
			} else {
				num = 9
			}
		case 7:
			num = 8
		}

		val += fmt.Sprintf("%v", num)
	}

	ans, _ := strconv.Atoi(val)
	return ans
}

func buildCounts(data string) (int, int) {
	easyCount := 0
	num := 0
	for _, entry := range strings.Split(data, "\n") {
		theory := make(map[int][]string)
		for i := 0; i <= 6; i++ {
			theory[i] = []string{"a", "b", "c", "d", "e", "f", "g"}
		}
		pieces := strings.Split(entry, "|")
		var sixes []string
		for _, word := range strings.Fields((strings.TrimSpace(pieces[0]))) {
			switch len(word) {
			case 2:
				// This is 1
				reduce(theory, []int{2, 5}, word)
			case 3:
				//This is 7
				reduce(theory, []int{0, 2, 5}, word)
			case 4:
				//This is 4
				reduce(theory, []int{1, 2, 3, 5}, word)
			case 5:
				// This is 2, 3, 5

			case 6:
				// This is 0,6,9
				sixes = append(sixes, word)
			case 7:
				//This is eight - we learn nothing here
			}
		}

		//Work on the sixes
		overlap := buildOverlap(sixes)
		for _, key := range []int{2, 3, 4} {
			theory[key] = remove(theory[key], overlap)
		}

		// Now tidy
		done := ""
		for _, chararr := range theory {
			if len(chararr) == 1 {
				done += chararr[0]
			}
		}

		for key, chararr := range theory {
			if len(chararr) > 1 {
				theory[key] = remove(theory[key], done)
				if len(theory[key]) > 1 {
					log.Fatalf("NEEDS MORE WORK")
				}
			}
		}

		num += resolve(invert(theory), strings.TrimSpace(pieces[1]))
	}

	for _, entry := range strings.Split(data, "\n") {
		pieces := strings.Split(entry, "|")
		for _, word := range strings.Fields((strings.TrimSpace(pieces[1]))) {
			switch len(word) {
			case 2, 3, 4, 7:
				easyCount++
			}
		}
	}

	return easyCount, num
}

func (s *Server) Solve2021day8part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-8.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	n1, _ := buildCounts(trimmed)
	return &pb.SolveResponse{Answer: int32(n1)}, nil
}

func (s *Server) Solve2021day8part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-8.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	_, n2 := buildCounts(trimmed)
	return &pb.SolveResponse{Answer: int32(n2)}, nil
}
