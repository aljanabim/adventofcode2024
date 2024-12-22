package day16

import (
	"fmt"
	"math"
	"slices"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Maze struct {
	Width  int
	Height int
	Start  *Node
	End    *Node
	Nodes  map[[2]int]*Node
}

type Node struct {
	Pos   [2]int
	Up    *Node
	Right *Node
	Down  *Node
	Left  *Node
}

type Heading int

const (
	UP Heading = iota + 1
	RIGHT
	DOWN
	LEFT
	MAX_HEADING
)

func createMaze(lines []string) Maze {
	maze := Maze{Height: len(lines) - 1, Width: len(lines[0]), Nodes: map[[2]int]*Node{}}
	// Create nodes in maze
	for row, line := range lines {
		for col, cell := range line {
			node := &Node{Pos: [2]int{row, col}}
			if cell != '#' {
				maze.Nodes[[2]int{row, col}] = node
			}
			if cell == 'S' {
				maze.Start = node
			} else if cell == 'E' {
				maze.End = node
			}
		}
	}

	// Connect nodes in maze
	for pos, node := range maze.Nodes {
		up := [2]int{pos[0] - 1, pos[1]}
		right := [2]int{pos[0], pos[1] + 1}
		down := [2]int{pos[0] + 1, pos[1]}
		left := [2]int{pos[0], pos[1] - 1}

		if n, ok := maze.Nodes[up]; ok {
			node.Up = n
		}
		if n, ok := maze.Nodes[right]; ok {
			node.Right = n
		}
		if n, ok := maze.Nodes[down]; ok {
			node.Down = n
		}
		if n, ok := maze.Nodes[left]; ok {
			node.Left = n
		}
	}
	return maze
}

/*
######
#...E#
##.#.#
##...#
##.#.#
##S..#
######
*/

func Argmin(list []float64, value float64) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

func Search(node *Node, prevNode *Node, heading Heading, maze *Maze, visited map[[2]int]bool, steps int) float64 {
	steps++
	visited[node.Pos] = true

	if steps >= maze.Width*maze.Height {
		return math.Inf(1)
	}

	// reaching a start node is penalized
	if node == maze.Start && prevNode != maze.Start {
		// fmt.Println("Hit START at", node.Pos)
		for k := range visited {
			visited[k] = false
		}
		return math.Inf(1)
	}

	// reaching end is rewarding
	if node == maze.End {
		// fmt.Println("Hit END", node.Pos)
		for k := range visited {
			visited[k] = false
		}
		return 0
	}

	// reaching own tail is penalized and not resetting in order to continue the current search
	if (node.Up == prevNode || node.Up == nil || visited[node.Up.Pos]) &&
		(node.Right == prevNode || node.Right == nil || visited[node.Right.Pos]) &&
		(node.Down == prevNode || node.Down == nil || visited[node.Down.Pos]) &&
		(node.Left == prevNode || node.Left == nil || visited[node.Left.Pos]) {
		// fmt.Println("Hit TAIL", node.Pos)
		return math.Inf(1)
	}

	// reaching a dead-end is penalized (needed in case we hit end and in next turn we hit dead end
	// then previous one using )
	// if (node.Up == prevNode || node.Up == nil) && (node.Down == prevNode || node.Down == nil) &&
	// 	(node.Left == prevNode || node.Left == nil) && (node.Right == prevNode || node.Right == nil) {
	// 	fmt.Println("Hit DEAD-END", node.Pos, "from", prevNode.Pos)
	// 	// for k := range visited {
	// 	// 	visited[k] = false
	// 	// }
	// 	return math.Inf(1)
	// }

	searched := []float64{}
	if node.Up != nil && node.Up != prevNode && !visited[node.Up.Pos] {
		cost := 1.0
		if heading != UP {
			cost = 1000
		}
		// fmt.Println("Going up from", node.Pos, cost)
		searched = append(searched, cost+Search(node.Up, node, UP, maze, visited, steps))
	}
	if node.Right != nil && node.Right != prevNode && !visited[node.Right.Pos] {
		cost := 1.0
		if heading != RIGHT {
			cost = 1000
		}
		// fmt.Println("Going right from", node.Pos, cost)
		searched = append(searched, cost+Search(node.Right, node, RIGHT, maze, visited, steps))
	}
	if node.Down != nil && node.Down != prevNode && !visited[node.Down.Pos] {
		cost := 1.0
		if heading != DOWN {
			cost = 1000
		}
		// fmt.Println("Going down from", node.Pos, cost)
		searched = append(searched, cost+Search(node.Down, node, DOWN, maze, visited, steps))
	}
	if node.Left != nil && node.Left != prevNode && !visited[node.Left.Pos] {
		cost := 1.0
		if heading != LEFT {
			cost = 1000
		}
		// fmt.Println("Going left from", node.Pos, cost)
		searched = append(searched, cost+Search(node.Left, node, LEFT, maze, visited, steps))
	}

	return slices.Min(searched)
}

func solvePart1(lines []string) int {
	maze := createMaze(lines)
	for pos, node := range maze.Nodes {
		fmt.Println(pos, node)
	}
	fmt.Println(maze.Start)
	return 0
}

func Solve() {
	lines, err := utils.ReadLines("day16/input_mini.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(16, 1, res)
}
