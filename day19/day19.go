package day19

import (
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

// returns slices of stripes and flags
func readInput(lines []string) ([]string, []string) {
	stripesRaw := strings.Split(lines[0], ",")
	stripes := make([]string, len(stripesRaw))
	for i, stripe := range stripesRaw {
		stripes[i] = strings.TrimSpace(stripe)
	}
	flags := lines[2:]
	return stripes, flags
}

type key struct {
	stripe string
	depth  int
}

func findCombos(depth int, flag string, stripes []string, cache map[key]int) int {
	if depth == len(flag) {
		return 1
	}
	count := 0
	for _, stripe := range stripes {
		if depth+len(stripe) <= len(flag) && flag[depth:depth+len(stripe)] == stripe {
			if val, ok := cache[key{stripe, depth + len(stripe)}]; ok {
				count += val
			} else {
				add := findCombos(depth+len(stripe), flag, stripes, cache)
				cache[key{stripe, depth + len(stripe)}] = add
				count += add
			}
		}
	}
	return count
}

func Solve() {
	lines, err := utils.ReadLines("day19/input.txt")
	if err != nil {
		panic(err)
	}
	stripes, flags := readInput(lines)
	countUnique := 0
	countTotal := 0
	defer utils.Duration(utils.Track("Solve Day 1"))
	for _, flag := range flags {
		count := findCombos(0, flag, stripes, map[key]int{})
		if count > 0 {
			countUnique++
			countTotal += count
		}
	}
	utils.PrintSolution(19, 1, countUnique)
	utils.PrintSolution(19, 2, countTotal)
}
