package day15

import (
	"fmt"
	"os"
	"strings"
)

type Entity int

const (
	FREE Entity = iota
	BOX
	WALL
)

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

var dirMap = map[Dir][2]int{
	UP:    {0, -1},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
	RIGHT: {1, 0},
}

type Robot struct {
	X     int
	Y     int
	moves []Dir
}

type Warehouse struct {
	Grid  [][]Entity
	Robot Robot
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
				row[col] = FREE
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

func readMoves(file string, robot *Robot) error {
	rawData, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	for _, f := range strings.Split(string(rawData), "") {
		switch f {
		case "^":
			robot.moves = append(robot.moves, UP)
		case "v":
			robot.moves = append(robot.moves, DOWN)
		case ">":
			robot.moves = append(robot.moves, RIGHT)
		case "<":
			robot.moves = append(robot.moves, LEFT)
		}
	}
	return nil
}

func computeGPSSum(warehouse *Warehouse) int {

	return 0
}

func Solve() {
	rawData, err := os.ReadFile("day15/map_mini.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawData), "\n")
	warehouse := readMap(lines)

	for _, row := range warehouse.Grid {
		fmt.Println(row)
	}
	err = readMoves("day15/moves_mini.txt", &warehouse.Robot)
	if err != nil {
		panic(err)
	}
	fmt.Println(warehouse.Robot)
}
