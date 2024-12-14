package day04

import (
	"fmt"
	"log"

	"github.com/aljanabim/adventofcode2024/utils"
)

/*
		XMAS
		||||
depth   0123
		||||
        SAMX

    MAS
    \|/
   B-X-S
	/|\
	MAB
*/

func boundaryCheck(fulltxt []string, row, col int) bool {
	if row < 0 || row >= len(fulltxt) {
		return false
	}
	if col < 0 || col >= len(fulltxt[0]) {
		return false
	}
	return true
}

func check(fulltxt []string, pattern string, direction [2]int, depth int, row, col int) (bool, error) {
	if depth > len(pattern)-1 {
		return false, fmt.Errorf("went too deep [%d] for pattern %s", depth, pattern)
	}
	if fulltxt[row][col] == pattern[depth] && depth == len(pattern)-1 {
		return true, nil
	} else if fulltxt[row][col] == pattern[depth] && depth < len(pattern)-1 {
		newRow := row + direction[0]
		newCol := col + direction[1]
		newDepth := depth + 1
		if boundaryCheck(fulltxt, newRow, newCol) {
			return check(fulltxt, pattern, direction, newDepth, newRow, newCol)
		}

	}
	return false, nil
}

func solvePart1() int {
	var directions map[string][2]int = map[string][2]int{
		"tl": {-1, -1},
		"t":  {-1, 0},
		"tr": {-1, 1},
		"l":  {0, -1},
		"r":  {0, 1},
		"bl": {1, -1},
		"b":  {1, 0},
		"br": {1, 1},
	}

	lines, err := utils.ReadLines("day04/input")
	if err != nil {
		panic(err)
	}

	counter := 0
	for row := range len(lines) {
		for col := range len(lines[row]) {
			pattern := ""
			if lines[row][col] == 'X' {
				pattern = "XMAS"
			} else if lines[row][col] == 'S' {
				pattern = "SAMX"
			}
			if len(pattern) > 0 {
				for _, v := range directions {
					found, err := check(lines, "XMAS", v, 0, row, col)
					if err != nil {
						panic(err)
					}
					if found {
						counter++
					}
				}
			}

		}
	}
	return counter
}

func solvePart2() int {
	lines, err := utils.ReadLines("day04/input")
	if err != nil {
		log.Fatal(err)
	}
	counter := 0
	for i := range len(lines) - 2 {
		r1 := lines[i]
		r2 := lines[i+1]
		r3 := lines[i+2]
		cols := min(len(r1), len(r2), len(r3))
		for j := range cols - 2 {
			if r2[j+1] == 'A' && (r1[j] == 'M' && r3[j+2] == 'S' || r1[j] == 'S' && r3[j+2] == 'M') && (r3[j] == 'M' && r1[j+2] == 'S' || r3[j] == 'S' && r1[j+2] == 'M') {
				counter++
			}
		}
	}
	return counter

}
func Solve() {
	res := solvePart1()
	utils.PrintSolution(4, 1, res)
	res = solvePart2()
	utils.PrintSolution(4, 2, res)
}
