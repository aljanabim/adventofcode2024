package day20

import (
	"fmt"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Entity int

const (
	FREE Entity = iota
	WALL
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
	MAX_DIR
)

var dir2Move = map[Direction][2]int{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

type Maze struct {
	Grid  [][]Entity
	Start [2]int
	End   [2]int
}

func buildMaze(lines []string) Maze {
	maze := Maze{
		Grid: make([][]Entity, len(lines)),
	}

	for row, line := range lines {
		gridRow := make([]Entity, len(line))
		for col, c := range line {
			switch c {
			case '.':
				gridRow[col] = FREE
			case 'S':
				gridRow[col] = FREE
				maze.Start = [2]int{row, col}
			case 'E':
				gridRow[col] = FREE
				maze.End = [2]int{row, col}
			case '#':
				gridRow[col] = WALL
			}
		}
		maze.Grid[row] = gridRow
	}
	return maze
}

func walkMaze(maze *Maze) [][2]int {
	step := maze.Start
	path := [][2]int{step}
	visited := map[[2]int]bool{step: true}
	for step != maze.End {
		for _, move := range dir2Move {
			newStep := [2]int{step[0] + move[0], step[1] + move[1]}
			if maze.Grid[newStep[0]][newStep[1]] == FREE && !visited[newStep] {
				visited[newStep] = true
				step = newStep
				break
			}
		}
		path = append(path, step)
	}
	return path
}

func getSavingsFreq(path [][2]int, maxCheat int, maze *Maze) map[int]int {
	defer utils.Duration(utils.Track(fmt.Sprintf("Get Savings Freq with maxCheat %d", maxCheat)))
	saving2Freq := map[int]int{} // maps number of pico seconds saved to frequency
	for i, step := range path {
		for j, cheatStep := range path[i+1:] {
			cheatDist := utils.Abs(step[0]-cheatStep[0]) + utils.Abs(step[1]-cheatStep[1])
			cheatSaving := j - cheatDist + 1
			if cheatDist <= maxCheat && cheatSaving > 0 {
				saving2Freq[cheatSaving] += 1
			}
		}
	}
	return saving2Freq
}

func solveParts(lines []string, saveLimit int, maxCheat int) int {
	maze := buildMaze(lines)
	path := walkMaze(&maze)
	savings2Freq := getSavingsFreq(path, maxCheat, &maze)
	count := 0
	for saving, freq := range savings2Freq {
		if saving >= saveLimit {
			count += freq
		}
	}
	return count
}

func Solve() {
	lines, err := utils.ReadLines("day20/input.txt")
	if err != nil {
		panic(err)
	}
	res := solveParts(lines, 100, 2)
	utils.PrintSolution(20, 1, res)
	res = solveParts(lines, 100, 20)
	utils.PrintSolution(20, 2, res)
}
