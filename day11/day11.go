package day11

import (
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func blink(stones []string) []string {
	// skipNext := false
	for i, stone := range stones {
		if stone == "0" {
			stones[i] = "1"
		} else if len(stone)%2 == 0 {
			stones[i] = stone[:len(stone)/2]
			rightNum, err := strconv.ParseInt(stone[len(stone)/2:], 10, 64)
			if err != nil {
				panic(err)
			}
			rightNumStr := strconv.Itoa(int(rightNum))
			stones = append(stones, rightNumStr)
		} else if (len(stone)+1)%2 == 0 {
			stoneNum, err := strconv.ParseInt(stone, 10, 64)
			if err != nil {
				panic(err)
			}
			stones[i] = strconv.Itoa(int(stoneNum * 2024))
		}
	}
	return stones
}

func solvePart1(stones []string, nBlinks int) int {
	for range nBlinks {
		stones = blink(stones)
	}
	return len(stones)
}

func solvePart2(stones []string, nBlinks int) int {
	// cache := map[string][]string{}
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

// func blink(stone string) []string {
// 	stones := []string{}
// 	if stone == "0" {
// 		stones = append(stones, "1")
// 	} else if len(stone)%2 == 0 {
// 		stones = append(stones, stone[:len(stone)/2])
// 		rightNum, err := strconv.ParseInt(stone[len(stone)/2:], 10, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		rightNumStr := strconv.Itoa(int(rightNum))
// 		stones = append(stones, rightNumStr)
// 	} else if (len(stone)+1)%2 == 0 {
// 		stoneNum, err := strconv.ParseInt(stone, 10, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		stones = append(stones, strconv.Itoa(int(stoneNum*2024)))
// 	}
// 	return stones
// }
