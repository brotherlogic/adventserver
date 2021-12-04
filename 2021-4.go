package main

import (
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func buildNumArr(str string) []int {
	var nums []int
	for _, st := range strings.Split(str, ",") {
		n, _ := strconv.Atoi(st)
		nums = append(nums, n)
	}

	return nums
}

type board struct {
	nums [][]int
	seen [][]bool
}

func (b *board) play(num int) {
	for i := range b.nums {
		for j := range b.nums[i] {
			if b.nums[i][j] == num {
				b.seen[i][j] = true
				return
			}
		}
	}
}

func (b *board) won() bool {
	for row := range b.nums {
		allseen := true
		for col := range b.nums {
			if b.seen[row][col] == false {
				allseen = false
			}
		}

		if allseen {
			return true
		}
	}

	for col := range b.nums {
		allseen := true
		for row := range b.nums {
			if b.seen[row][col] == false {
				allseen = false
			}
		}

		if allseen {
			return true
		}
	}

	return false
}

func (b *board) score() int {
	total := 0
	for i := range b.nums {
		for j := range b.nums {
			if !b.seen[i][j] {
				total += b.nums[i][j]
			}
		}
	}

	return total
}

func buildBoard(strs []string) *board {
	board := &board{}
	for _, line := range strs {
		var nums []int
		var seens []bool
		bits := strings.Fields(line)
		for _, bit := range bits {
			num, _ := strconv.Atoi(bit)
			nums = append(nums, num)
			seens = append(seens, false)
		}

		board.nums = append(board.nums, nums)
		board.seen = append(board.seen, seens)
	}

	return board
}

func runBingo(data string) int {
	bits := strings.Split(data, "\n")
	nums := buildNumArr(bits[0])

	var boards []*board
	for i := 2; i < len(bits); i += 6 {
		boards = append(boards, buildBoard(bits[i:i+5]))
	}

	for _, num := range nums {
		for _, board := range boards {
			board.play(num)
			if board.won() {
				return board.score() * num
			}
		}
	}

	return 0
}

func runBingoLast(data string) int {
	bits := strings.Split(data, "\n")
	nums := buildNumArr(bits[0])

	var boards []*board
	for i := 2; i < len(bits); i += 6 {
		boards = append(boards, buildBoard(bits[i:i+5]))
	}

	for _, num := range nums {
		var nboards []*board
		for _, board := range boards {
			board.play(num)
			if board.won() {
				if len(boards) == 1 {
					return board.score() * num
				}
			} else {
				nboards = append(nboards, board)
			}
		}
		boards = nboards
	}

	return 0
}

func (s *Server) Solve2021day4part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-4.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runBingo(data))}, nil
}

func (s *Server) Solve2021day4part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-4.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(runBingoLast(data))}, nil
}
