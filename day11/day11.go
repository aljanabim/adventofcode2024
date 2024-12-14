package day11

import (
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func blink(stones []string) []string {
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
	defer utils.Duration(utils.Track("Part1"))
	for range nBlinks {
		stones = blink(stones)
	}
	return len(stones)
}

type Record struct {
	string
	int
}
type Cache map[Record]int

func explore(stones []string, blinks int, cache Cache) int {
	if blinks == 0 {
		return len(stones)
	}
	count := 0
	for _, stone := range stones {
		if val, ok := cache[Record{stone, blinks}]; ok {
			count += val
			continue
		}
		subTot := explore(blink([]string{stone}), blinks-1, cache)
		cache[Record{stone, blinks}] = subTot
		count += subTot
	}
	return count
}

func solvePart2(stones []string, blinks int) int {
	defer utils.Duration(utils.Track("Part2"))

	var cache = make(Cache)
	tot := 0
	tot += explore(stones, blinks, cache)
	return tot
}

func Solve() {
	file, err := os.ReadFile("day11/input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(file)), " ")

	res := solvePart1(input, 30)
	utils.PrintSolution(11, 1, res)
	input = strings.Split(strings.TrimSpace(string(file)), " ")
	res = solvePart2(input, 75)
	utils.PrintSolution(11, 2, res)
}
