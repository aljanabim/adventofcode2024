package day20

import (
	"math"

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

func getSavingsFreq(path [][2]int, maxCheat int) map[int]int {
	saving2Freq := map[int]int{} // maps number of pico seconds saved to frequency
	for i, step := range path {
		// fmt.Println(i, step)
		for depth := 2; depth <= maxCheat; depth++ {
			upPos := [2]int{step[0] + depth*dir2Move[UP][0], step[1] + depth*dir2Move[UP][1]}
			rightPos := [2]int{step[0] + depth*dir2Move[RIGHT][0], step[1] + depth*dir2Move[RIGHT][1]}
			downPos := [2]int{step[0] + depth*dir2Move[DOWN][0], step[1] + depth*dir2Move[DOWN][1]}
			leftPos := [2]int{step[0] + depth*dir2Move[LEFT][0], step[1] + depth*dir2Move[LEFT][1]}

			for j, cheatStep := range path[i+1:] {
				if cheatStep == upPos || cheatStep == rightPos || cheatStep == downPos || cheatStep == leftPos {
					if j-1 > 0 {
						saving2Freq[j-1]++
					}
				}
			}
			// fmt.Println("look to sides", leftPos, rightPos)
		}
	}
	return saving2Freq
}

/*
.....3.....
....323....
...3...3...
..32.O.23..  O @ (3,5)
...3...3...
....323....
.....3.....
*/

func getSavingsFreq2(path [][2]int, maxCheat int, maze *Maze) map[int]int {
	saving2Freq := map[int]int{} // maps number of pico seconds saved to frequency
	for i, step := range path {
		reachablePoints := map[[2]int]int{}
		rowLowerLim := int(math.Max(float64(step[0]-maxCheat), 0))
		rowUpperLim := int(math.Min(float64(step[0]+maxCheat), float64(len(maze.Grid))-1))
		colLowerLim := int(math.Max(float64(step[1]-maxCheat), 0))
		colUpperLim := int(math.Min(float64(step[1]+maxCheat), float64(len(maze.Grid[0]))-1))

		for row := rowLowerLim; row <= rowUpperLim; row++ {
			for col := colLowerLim; col <= colUpperLim; col++ {
				dist := utils.Abs(step[0]-row) + utils.Abs(step[1]-col)
				if 1 <= dist && dist <= maxCheat && maze.Grid[row][col] == FREE {
					reachablePoints[[2]int{row, col}] = dist
				}
			}
		}

		for j, cheatStep := range path[i+1:] {
			cheatDist := reachablePoints[cheatStep]
			cheatSaving := j - cheatDist + 1
			if cheatDist > 0 && cheatSaving > 0 {
				saving2Freq[cheatSaving] += 1
			}
		}
	}
	return saving2Freq
}

func solvePart1(lines []string, saveLimit int) int {
	defer utils.Duration(utils.Track("Part1"))
	maze := buildMaze(lines)
	path := walkMaze(&maze)
	savings2Freq := getSavingsFreq(path, 2)
	count := 0
	for saving, freq := range savings2Freq {
		if saving >= saveLimit {
			count += freq
		}
	}
	return count
}

func solvePart2(lines []string, saveLimit int, maxCheat int) int {
	defer utils.Duration(utils.Track("Part2"))
	maze := buildMaze(lines)
	path := walkMaze(&maze)
	savings2Freq := getSavingsFreq2(path, maxCheat, &maze)
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

	res := solvePart1(lines, 100)
	utils.PrintSolution(20, 1, res)
	res = solvePart2(lines, 100, 2)
	res = solvePart2(lines, 100, 20)
	utils.PrintSolution(20, 2, res)

}
