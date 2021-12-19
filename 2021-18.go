package main

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type snnum struct {
	val    int
	left   *snnum
	right  *snnum
	parent *snnum
}

func explode(num *snnum) *snnum {
	runExplode(num, nil, 0)

	return num
}

func fullExplode(num *snnum) *snnum {
	for runExplode(num, nil, 0) {
		//Pass
	}

	return num
}

func reduceNum(num *snnum) *snnum {
	for {
		fullExplode(num)
		if runSplit(num) {
			split(num)
		} else {
			break
		}
	}

	return num
}

func add(num1, num2 *snnum) *snnum {
	val := &snnum{left: num1, right: num2, val: -1}
	num1.parent = val
	num2.parent = val
	return val
}

func runSum(data string) *snnum {
	elems := strings.Split(strings.TrimSpace(data), "\n")
	num := parseNum(strings.TrimSpace(elems[0]))
	for i := 1; i < len(elems); i++ {
		nnum := reduceNum(add(num, parseNum(strings.TrimSpace(elems[i]))))
		num = nnum
	}

	return num
}

func getLeft(num *snnum) *snnum {
	root := num
	for root.parent != nil {
		root = root.parent
	}

	list := listify(root)
	found := false
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == num {
			found = true
		} else if found && (list[i].val >= 0) {
			return list[i]
		}
	}

	return nil
}

func runSplit(num *snnum) bool {
	if num == nil {
		return false
	}

	if num.val > 9 {
		split(num)
		return true
	}

	return runSplit(num.left) || runSplit(num.right)
}

func split(num *snnum) *snnum {
	if num.val > 9 {
		adjust := 0
		if num.val%2 != 0 {
			adjust++
		}
		num.left = &snnum{parent: num, val: num.val / 2}
		num.right = &snnum{parent: num, val: num.val/2 + adjust}
		num.val = -1
	}
	return num
}

func getRight(num *snnum) *snnum {
	root := num
	for root.parent != nil {
		root = root.parent
	}

	list := listify(root)
	found := false
	for i := 0; i < len(list); i++ {
		if list[i] == num {
			found = true
		} else if found && (list[i].val >= 0) {
			return list[i]
		}
	}

	return nil
}

func runExplode(num *snnum, parent *snnum, depth int) bool {
	if num == nil {
		return false
	}

	if depth >= 4 && num.left != nil && num.right != nil && num.left.val >= 0 && num.right.val >= 0 {
		// Run the explode
		lnum := getLeft(num.left)
		rnum := getRight(num.right)

		if lnum != nil {
			lnum.val += num.left.val
		}
		if rnum != nil {
			rnum.val += num.right.val
		}

		num.left = nil
		num.right = nil
		num.val = 0

		return true
	}

	// Search the left tree
	found := runExplode(num.left, num, depth+1)

	if !found {
		return runExplode(num.right, num, depth+1)
	}

	return true
}

func listify(num *snnum) []*snnum {
	var str []*snnum
	if num.left != nil {
		str = append(str, listify(num.left)...)
	}

	str = append(str, num)

	if num.right != nil {
		str = append(str, listify(num.right)...)
	}

	return str
}

func printNum(num *snnum) string {
	if num == nil {
		return "nil"
	}
	if num.val >= 0 {
		return fmt.Sprintf("%v", num.val)
	}
	str := "["
	if num.left != nil {
		str += printNum(num.left)
	}
	str += ","

	if num.right != nil {
		str += printNum(num.right)
	}

	str += "]"
	return str
}

func findLeftBalance(str string) int {
	count := 0
	for i := range str {
		if string(str[i]) == "[" {
			count++
		} else if string(str[i]) == "]" {
			count--
		}

		if count == 0 {
			return i
		}
	}

	return -1
}

func findRightBalance(str string) int {
	count := 0
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "]" {
			count++
		} else if string(str[i]) == "[" {
			count--
		}

		if count == 0 {
			return i
		}
	}

	return -1
}

func magnitude(num *snnum) int {
	if num.val >= 0 {
		return num.val
	}
	return 3*magnitude(num.left) + 2*magnitude(num.right)
}

func parseNum(str string) *snnum {
	val, _ := parseNumInt(str, nil)
	return val
}

func bestSum(data string) int {
	elems := strings.Split(strings.TrimSpace(data), "\n")
	best := 0
	for i := 0; i < len(elems); i++ {
		for j := i + 1; j < len(elems); j++ {
			b1 := magnitude(reduceNum(add(parseNum(strings.TrimSpace(elems[i])), parseNum(strings.TrimSpace(elems[j])))))
			if b1 > best {
				best = b1
			}
			b1 = magnitude(reduceNum(add(parseNum(strings.TrimSpace(elems[j])), parseNum(strings.TrimSpace(elems[i])))))
			if b1 > best {
				best = b1
			}
		}
	}

	return best
}

func parseNumInt(str string, curr *snnum) (*snnum, string) {
	ret := &snnum{parent: curr, val: -1}

	switch string(str[0]) {
	case "[":
		val, rem := parseNumInt(str[1:], &snnum{parent: ret, val: -1})
		ret.left = val
		ret.right, rem = parseNumInt(rem[1:], &snnum{parent: ret, val: -1})
		return ret, rem[1:]
	case "]":
		return ret, str[1:]
	default:
		val, _ := strconv.Atoi(string(str[0]))
		ret.val = val
		return ret, str[1:]
	}
}

func (s *Server) Solve2021day18part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-18.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(magnitude(runSum(trimmed)))}, nil
}

func (s *Server) Solve2021day18part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-18.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(bestSum(trimmed))}, nil
}
