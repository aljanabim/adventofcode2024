package day11

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func blink(stones []string) []string {
	skipNext := false
	i := 0
	for i < len(stones) {
		stone := stones[i]
		if skipNext {
			i++
			skipNext = false
			continue
		}
		if stone == "0" {
			stones[i] = "1"
		} else if len(stone)%2 == 0 {
			stones[i] = stone[:len(stone)/2]
			rightNum, err := strconv.ParseInt(stone[len(stone)/2:], 10, 64)
			if err != nil {
				panic(err)
			}
			rightNumStr := strconv.Itoa(int(rightNum))
			stones = slices.Insert(stones, i+1, rightNumStr)
			skipNext = true
		} else if (len(stone)+1)%2 == 0 {
			stoneNum, err := strconv.ParseInt(stone, 10, 64)
			if err != nil {
				panic(err)
			}
			stones[i] = strconv.Itoa(int(stoneNum * 2024))
		}
		i++
	}
	return stones
}

func solvePart1(stones []string, nBlinks int) int {
	for range nBlinks {
		stones = blink(stones)
	}
	return len(stones)
}

func Solve() {
	file, err := os.ReadFile("day11/input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(file)), " ")

	res := solvePart1(input, 25)
	utils.PrintSolution(11, 1, res)
}
