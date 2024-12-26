package day20

import (
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

type Step struct {
	pos [2]int
	dir Direction
}

func walkMaze(maze *Maze) []Step {
	step := Step{pos: maze.Start}
	path := []Step{step}
	visited := map[[2]int]bool{step.pos: true}
	for step.pos != maze.End {
		for dir, move := range dir2Move {
			newStep := Step{pos: [2]int{step.pos[0] + move[0], step.pos[1] + move[1]}, dir: dir}
			if maze.Grid[newStep.pos[0]][newStep.pos[1]] == FREE && !visited[newStep.pos] {
				visited[newStep.pos] = true
				step = newStep
				break
			}
		}
		path = append(path, step)
	}
	// shift dir one step back
	for i := range len(path) - 1 {
		path[i].dir = path[i+1].dir
	}
	return path
}

func getSavingsFreq(path []Step) map[int]int {
	saving2Freq := map[int]int{} // maps number of pico seconds saved to frequency
	for i, step := range path {
		// fmt.Println(i, step)
		upPos := [2]int{step.pos[0] + 2*dir2Move[UP][0], step.pos[1] + 2*dir2Move[UP][1]}
		rightPos := [2]int{step.pos[0] + 2*dir2Move[RIGHT][0], step.pos[1] + 2*dir2Move[RIGHT][1]}
		downPos := [2]int{step.pos[0] + 2*dir2Move[DOWN][0], step.pos[1] + 2*dir2Move[DOWN][1]}
		leftPos := [2]int{step.pos[0] + 2*dir2Move[LEFT][0], step.pos[1] + 2*dir2Move[LEFT][1]}

		for j, cheatStep := range path[i+1:] {
			if cheatStep.pos == upPos || cheatStep.pos == rightPos || cheatStep.pos == downPos || cheatStep.pos == leftPos {
				if j-1 > 0 {
					saving2Freq[j-1]++
					// fmt.Println("Cheat at", j+i+1, cheatStep.pos, "saving", j-1)
				}
			}
		}
		// fmt.Println("look to sides", leftPos, rightPos)
	}
	return saving2Freq
}

func solvePart1(lines []string) int {
	maze := buildMaze(lines)
	path := walkMaze(&maze)
	savings2Freq := getSavingsFreq(path)
	count := 0
	for saving, freq := range savings2Freq {
		if saving >= 100 {
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
	res := solvePart1(lines)
	utils.PrintSolution(20, 1, res)

}
