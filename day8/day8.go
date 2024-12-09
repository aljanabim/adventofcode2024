package day8

import (
	"github.com/aljanabim/adventofcode2024/utils"
)

func solvePart1(lines []string) int {
	antennas := map[rune][][2]int{}
	antinodes := map[[2]int]int{}
	rows := len(lines)
	cols := len(lines[0])
	for row, line := range lines {
		for col, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], [2]int{row, col})
			}
		}
	}
	for _, pos := range antennas {
		for i, p1 := range pos {
			for _, p2 := range pos[i+1:] {
				dx := p2[1] - p1[1]
				dy := p2[0] - p1[0]
				a1 := [2]int{p1[0] - dy, p1[1] - dx}
				a2 := [2]int{p2[0] + dy, p2[1] + dx}
				if utils.Inside(a1, rows, cols) {
					antinodes[a1]++
				}
				if utils.Inside(a2, rows, cols) {
					antinodes[a2]++
				}
			}
		}
	}
	return len(antinodes)
}

func solvePart2(lines []string) int {
	antennas := map[rune][][2]int{}
	antinodes := map[[2]int]int{}
	rows := len(lines)
	cols := len(lines[0])
	for row, line := range lines {
		for col, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], [2]int{row, col})
			}
		}
	}
	for _, pos := range antennas {
		for i, p1 := range pos {
			antinodes[p1]++
			for _, p2 := range pos[i+1:] {
				antinodes[p2]++

				dx := p2[1] - p1[1]
				dy := p2[0] - p1[0]

				xRepeats := cols / utils.Abs(dx)
				yRepeats := rows / utils.Abs(dx)
				repeats := max(xRepeats, yRepeats)
				for r := range repeats {
					a1 := [2]int{p1[0] - dy*r, p1[1] - dx*r}
					a2 := [2]int{p2[0] + dy*r, p2[1] + dx*r}
					if utils.Inside(a1, rows, cols) {
						antinodes[a1]++
					}
					if utils.Inside(a2, rows, cols) {
						antinodes[a2]++
					}
				}
				// fmt.Println("combo", p1, "with", p2, "repeat", repeats)
			}
		}
	}
	return len(antinodes)
}

func Solve() {
	lines, err := utils.ReadLines("day8/input")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(8, 1, res)
	res = solvePart2(lines)
	utils.PrintSolution(8, 2, res)
}
