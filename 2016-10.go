package main

import (
	"strconv"
	"strings"
)

type bot struct {
	low, high int
	nums      []int
	comps     [][]int
}

func (b bot) comp(val1, val2 int) bool {
	for _, c := range b.comps {
		if c[0] == val1 && c[1] == val2 || c[0] == val2 && c[1] == val1 {
			return true
		}
	}
	return false
}

func runBotProgram(data string) []*bot {
	bots := make(map[int]*bot)
	output := make(map[int]int)

	for _, line := range strings.Split(data, "\n") {
		fields := strings.Fields(line)
		if fields[0] == "value" {
			num, _ := strconv.ParseInt(fields[1], 10, 32)
			bnum, _ := strconv.ParseInt(fields[5], 10, 32)
			b, ok := bots[int(bnum)]
			if !ok {
				b = &bot{}
				bots[int(bnum)] = b
			}
			b.nums = append(b.nums, int(num))

		} else if fields[0] == "bot" {
			bnum, _ := strconv.ParseInt(fields[1], 10, 32)
			lnum, _ := strconv.ParseInt(fields[6], 10, 32)
			hnum, _ := strconv.ParseInt(fields[11], 10, 32)

			b, ok := bots[int(bnum)]
			if !ok {
				b = &bot{}
				bots[int(bnum)] = b
			}

			if fields[5] == "bot" {
				b.low = int(lnum)
			} else {
				b.low = -int(lnum)
			}

			if fields[10] == "bot" {
				b.high = int(hnum)
			} else {
				b.high = -int(hnum)
			}
		}
	}

	for {
		found := false
		for _, bot := range bots {
			if len(bot.nums) == 2 {
				found = true
				low := bot.nums[0]
				high := bot.nums[1]

				if bot.nums[0] > bot.nums[1] {
					high = bot.nums[0]
					low = bot.nums[1]
				}

				if bot.low > 0 {
					bl := bots[bot.low]
					bl.nums = append(bl.nums, low)
				} else {
					output[-bot.low] = low
				}
				if bot.high > 0 {
					bh := bots[bot.high]
					bh.nums = append(bh.nums, high)
				} else {
					output[-bot.high] = high
				}

				bot.comps = append(bot.comps, []int{low, high})
				bot.nums = []int{}

			}
		}
		if !found {
			break
		}
	}

	var botList []*bot
	for _, bot := range bots {
		botList = append(botList, bot)
	}
	return botList
}
