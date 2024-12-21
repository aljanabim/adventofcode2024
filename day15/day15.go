package day15

import (
	"os"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Entity int

const (
	FREE Entity = iota
	BOX
	WALL
	ROBOT
)

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

var dirMap = map[Dir][2]int{
	UP:    {-1, 0},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
	RIGHT: {0, 1},
}

type Robot struct {
	X int
	Y int
}

func (r *Robot) Move(dir Dir) {
	r.X += dirMap[dir][1]
	r.Y += dirMap[dir][0]
}

type Warehouse struct {
	Grid  [][]Entity
	Robot Robot
}

func (w *Warehouse) Step(moves []Dir) {
	for _, move := range moves {
		boxesToMove := [][2]int{}
		col := w.Robot.X
		row := w.Robot.Y
		moveRobot := false
		// Get boxes to move and whether to move robot
		for {
			row += dirMap[move][0]
			col += dirMap[move][1]
			entity := w.Grid[row][col]
			if entity == WALL {
				boxesToMove = [][2]int{}
				moveRobot = false
				break
			} else if entity == BOX {
				boxesToMove = append(boxesToMove, [2]int{row, col})
			} else if entity == FREE {
				moveRobot = true
				break
			}
		}
		// Move boxes
		for _, box := range boxesToMove {
			movedBox := [2]int{box[0] + dirMap[move][0], box[1] + dirMap[move][1]}
			w.Grid[movedBox[0]][movedBox[1]] = BOX
		}
		// Move robot
		if moveRobot {
			w.Grid[w.Robot.Y][w.Robot.X] = FREE
			w.Robot.Move(move)
			w.Grid[w.Robot.Y][w.Robot.X] = ROBOT
		}
	}
}

func readMap(lines []string) Warehouse {
	warehouse := Warehouse{
		Grid:  make([][]Entity, len(lines)),
		Robot: Robot{},
	}

	for i, line := range lines {
		row := make([]Entity, len(line))
		for col, char := range line {
			switch char {
			case '#':
				row[col] = WALL
			case 'O':
				row[col] = BOX
			case '@':
				row[col] = ROBOT
				warehouse.Robot.X = col
				warehouse.Robot.Y = i
			default:
				row[col] = FREE
			}
		}
		warehouse.Grid[i] = row
	}

	return warehouse
}

func readMoves(file string) ([]Dir, error) {
	rawData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	line := strings.Split(string(rawData), "")
	moves := make([]Dir, len(line))

	for i, f := range line {
		switch f {
		case "^":
			moves[i] = UP
		case "v":
			moves[i] = DOWN
		case ">":
			moves[i] = RIGHT
		case "<":
			moves[i] = LEFT
		}
	}
	return moves, nil
}

func computeGPSSum(warehouse *Warehouse) int {
	sum := 0
	for y, row := range warehouse.Grid {
		for x, e := range row {
			if e == BOX {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func solvePart1() int {
	rawData, err := os.ReadFile("day15/map.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawData), "\n")
	warehouse := readMap(lines)

	moves, err := readMoves("day15/moves.txt")
	if err != nil {
		panic(err)
	}

	warehouse.Step(moves)
	res := computeGPSSum(&warehouse)
	return res
}
func Solve() {
	res := solvePart1()
	utils.PrintSolution(15, 1, res)
}
