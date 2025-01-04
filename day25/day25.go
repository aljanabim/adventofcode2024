package day25

import (
	"os"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

// returns locks and keys
func parseInput(file string) ([][5]int, [][5]int, error) {
	keys := [][5]int{}
	locks := [][5]int{}
	rawBytes, err := os.ReadFile(file)
	if err != nil {
		return locks, keys, err
	}

	for _, item := range strings.Split(string(rawBytes), "\n\n") {
		lines := strings.Split(item, "\n")
		heights := [5]int{}
		if lines[0] == "#####" { // lock
			for _, row := range lines[1:] {
				for i, col := range row {
					if col == '#' {
						heights[i]++
					}
				}
			}
			locks = append(locks, heights)
		}
		if lines[len(lines)-1] == "#####" { // lock
			for _, row := range lines[:len(lines)-1] {
				for i, col := range row {
					if col == '#' {
						heights[i]++
					}
				}
			}
			keys = append(keys, heights)
		}
	}

	return locks, keys, nil
}

func checkOverlap(lock, key [5]int) bool {
	res := true
	for i := range len(lock) {
		res = res && lock[i]+key[i] <= len(lock)
	}
	return res
}

func solvePart1(locks, keys [][5]int) int {
	defer utils.Duration(utils.Track("Part 1"))
	res := 0
	for _, lock := range locks {
		for _, key := range keys {
			if checkOverlap(lock, key) {
				res++
			}
		}
	}
	return res
}

func Solve() {
	locks, keys, err := parseInput("day25/input.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(locks, keys)
	utils.PrintSolution(25, 1, res)
}
