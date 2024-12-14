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

var cache = make(Cache)

func explore(record Record) {
	if _, ok := cache[record]; ok {
		// fmt.Println("Found cached", record)
		return
	}
	// fmt.Println("Exploring", record)
	res := []string{record.string}
	for nBlink := range record.int {
		res = blink(res)
		cache[Record{record.string, nBlink + 1}] = len(res)
		// fmt.Println("Caching (", record.string, nBlink+1, "): ", len(res))
		for _, stone := range res {
			if record.int-nBlink > 1 {
				newRecord := Record{stone, record.int - nBlink - 1}
				explore(newRecord)
			}
		}
	}
}

func solvePart2(stones []string, blinks int) int {
	defer utils.Duration(utils.Track("Part2"))

	for _, stone := range stones {
		record := Record{stone, blinks}
		// fmt.Println("Part2 stone", stone)
		explore(record)
	}
	tot := 0
	for _, stone := range stones {
		tot += cache[Record{stone, blinks}]
	}
	return tot
}

func Solve() {
	file, err := os.ReadFile("day11/input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(file)), " ")

	res := solvePart1(input, 20)
	utils.PrintSolution(11, 1, res)
	input = strings.Split(strings.TrimSpace(string(file)), " ")
	res = solvePart2(input, 20)
	utils.PrintSolution(11, 2, res)
}
