package day15

import (
	"bufio"
	"fmt"
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
	BOX_LEFT
	BOX_RIGHT
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
		boxesToMoveExtended := [][2]int{}
		col := w.Robot.X
		row := w.Robot.Y
		moveRobot := false
		cols := map[int]bool{col: true}
		// Get boxes to move and whether to move robot
	outerloop:
		for {
			if move == RIGHT || move == LEFT {
				col += dirMap[move][1]
				entity := w.Grid[row][col]
				if entity == WALL {
					boxesToMove = [][2]int{}
					boxesToMoveExtended = [][2]int{}
					moveRobot = false
					break
				} else if entity == BOX {
					boxesToMove = append(boxesToMove, [2]int{row, col})
				} else if entity == FREE {
					moveRobot = true
					break
				} else if entity == BOX_LEFT {
					boxesToMoveExtended = append(boxesToMoveExtended, [2]int{row, col})
				} else if entity == BOX_RIGHT {
					continue
				}
			} else {
				row += dirMap[move][0]
				// add cols to check
				for col := range cols {
					entity := w.Grid[row][col]
					if entity == BOX_LEFT {
						cols[col+1] = true
					} else if entity == BOX_RIGHT {
						cols[col-1] = true
					} else if entity == FREE {
						delete(cols, col)
					}

				}
				allFree := true
				for col := range cols {
					entity := w.Grid[row][col]
					if entity != FREE {
						allFree = false
					}
					if entity == WALL {
						boxesToMove = [][2]int{}
						boxesToMoveExtended = [][2]int{}
						moveRobot = false
						break outerloop
					} else if entity == BOX {
						boxesToMove = append(boxesToMove, [2]int{row, col})
					} else if entity == BOX_LEFT {
						boxesToMoveExtended = append(boxesToMoveExtended, [2]int{row, col})
					}
				}
				if allFree {
					moveRobot = true
					break
				}

			}
		}
		// Move boxes
		for _, box := range boxesToMove {
			w.Grid[box[0]][box[1]] = FREE
		}
		for _, box := range boxesToMove {
			movedBox := [2]int{box[0] + dirMap[move][0], box[1] + dirMap[move][1]}
			w.Grid[movedBox[0]][movedBox[1]] = BOX
		}
		// Move boxes extended
		for _, box := range boxesToMoveExtended {
			w.Grid[box[0]][box[1]] = FREE
			w.Grid[box[0]][box[1]+1] = FREE
		}
		for _, box := range boxesToMoveExtended {
			movedBoxLeft := [2]int{box[0] + dirMap[move][0], box[1] + dirMap[move][1]}
			movedBoxRight := [2]int{movedBoxLeft[0], movedBoxLeft[1] + 1}
			w.Grid[movedBoxLeft[0]][movedBoxLeft[1]] = BOX_LEFT
			w.Grid[movedBoxRight[0]][movedBoxRight[1]] = BOX_RIGHT
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
			case '[':
				row[col] = BOX_LEFT
			case ']':
				row[col] = BOX_RIGHT
			default:
				row[col] = FREE
			}
		}
		warehouse.Grid[i] = row
	}

	return warehouse
}

func readMapExtended(lines []string) Warehouse {
	warehouse := Warehouse{
		Grid:  make([][]Entity, len(lines)),
		Robot: Robot{},
	}

	for i, line := range lines {
		row := make([]Entity, len(line)*2)
		for col, char := range line {
			switch char {
			case '#':
				row[2*col] = WALL
				row[2*col+1] = WALL
			case 'O':
				row[2*col] = BOX_LEFT
				row[2*col+1] = BOX_RIGHT
			case '@':
				row[2*col] = ROBOT
				warehouse.Robot.X = 2 * col
				warehouse.Robot.Y = i
				row[2*col+1] = FREE
			default:
				row[2*col] = FREE
				row[2*col+1] = FREE
			}
		}
		warehouse.Grid[i] = row
	}

	return warehouse
}

func printMap(grid [][]Entity) string {
	s := strings.Builder{}
	for i, row := range grid {
		for _, entity := range row {
			switch entity {
			case FREE:
				s.WriteRune('.')
			case BOX:
				s.WriteRune('O')
			case BOX_LEFT:
				s.WriteRune('[')
			case BOX_RIGHT:
				s.WriteRune(']')
			case ROBOT:
				s.WriteRune('@')
			case WALL:
				s.WriteRune('#')

			}
		}
		if i < len(grid)-1 {
			s.WriteString("\n")
		}
	}
	return s.String()
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
			if e == BOX || e == BOX_LEFT {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func solvePart1(map_path, moves_path string) int {
	rawData, err := os.ReadFile(map_path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawData), "\n")
	warehouse := readMap(lines)

	moves, err := readMoves(moves_path)
	if err != nil {
		panic(err)
	}

	warehouse.Step(moves)
	res := computeGPSSum(&warehouse)
	return res
}

func solvePart2(map_path, moves_path string, interactive bool) int {
	rawData, err := os.ReadFile(map_path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawData), "\n")
	warehouse := readMapExtended(lines)

	moves, err := readMoves(moves_path)
	if err != nil {
		panic(err)
	}

	if interactive {
		for _, m := range moves {
			warehouse.Step([]Dir{m})
			switch m {
			case LEFT:
				fmt.Println("<")
			case RIGHT:
				fmt.Println(">")
			case UP:
				fmt.Println("^")
			case DOWN:
				fmt.Println("v")
			}
			fmt.Print(printMap(warehouse.Grid))
			reader := bufio.NewReader(os.Stdin)
			_, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
		}
	} else {

		warehouse.Step(moves)
	}

	res := computeGPSSum(&warehouse)
	return res
}

func Solve() {
	res := solvePart1("day15/map.txt", "day15/moves.txt")
	utils.PrintSolution(15, 1, res)
	res = solvePart2("day15/map.txt", "day15/moves.txt", false)
	utils.PrintSolution(15, 2, res)
}
