package day10

import (
	"github.com/aljanabim/adventofcode2024/utils"
)

type Node struct {
	height  int
	visited bool
	up      *Node
	down    *Node
	left    *Node
	right   *Node
}

func traverse(n *Node, score int, startNode *Node) int {
	n.visited = true
	if n.height == 9 {
		if startNode != nil {
			resetVisited(startNode)
		}
		return score + 1
	}

	score1, score2, score3, score4 := 0, 0, 0, 0
	if n.up != nil && !n.up.visited {
		score1 = traverse(n.up, score, startNode)
	}
	if n.down != nil && !n.down.visited {
		score2 = traverse(n.down, score, startNode)
	}
	if n.left != nil && !n.left.visited {
		score3 = traverse(n.left, score, startNode)
	}
	if n.right != nil && !n.right.visited {
		score4 = traverse(n.right, score, startNode)
	}
	return score1 + score2 + score3 + score4
}

func resetVisited(n *Node) {
	n.visited = false

	if n.up != nil && n.up != n {
		resetVisited(n.up)
	}
	if n.down != nil && n.down != n {
		resetVisited(n.down)
	}
	if n.left != nil && n.left != n {
		resetVisited(n.left)
	}
	if n.right != nil && n.right != n {
		resetVisited(n.right)
	}

}

func createGrid(lines []string) [][]*Node {
	grid := make([][]*Node, len(lines))
	for row, line := range lines {
		r := []*Node{}
		for _, c := range line {
			r = append(r, &Node{height: int(c - '0')})
		}
		grid[row] = r
	}
	return grid
}

func updateGrid(grid [][]*Node) []*Node {
	trailStart := []*Node{}
	for i, row := range grid {
		for j, node := range row {
			if j > 0 { // left neighbor
				leftNode := grid[i][j-1]
				if leftNode.height-node.height == 1 {
					node.left = leftNode
				}
			}
			if j < len(row)-1 { // right neighbor
				rightNode := grid[i][j+1]
				if rightNode.height-node.height == 1 {
					node.right = rightNode
				}
			}

			if i > 0 { // up neighbor
				upNode := grid[i-1][j]
				if upNode.height-node.height == 1 {
					node.up = upNode
				}
			}
			if i < len(grid)-1 { // down neighbor
				downNode := grid[i+1][j]
				if downNode.height-node.height == 1 {
					node.down = downNode
				}
			}

			if node.height == 0 {
				trailStart = append(trailStart, node)
			}
		}
	}
	return trailStart
}

func solvePart1(lines []string) int {
	defer utils.Duration(utils.Track("Part 1"))
	grid := createGrid(lines)
	trailHeads := updateGrid(grid)
	score := 0
	for _, n := range trailHeads {
		headScore := traverse(n, 0, nil)
		resetVisited(n)
		score += headScore
	}
	return score

}

func solvePart2(lines []string) int {
	defer utils.Duration(utils.Track("Part 2"))
	grid := createGrid(lines)
	trailHeads := updateGrid(grid)
	score := 0
	for _, n := range trailHeads {
		headScore := traverse(n, 0, n)
		resetVisited(n)
		score += headScore
	}
	return score
}

func Solve() {
	lines, err := utils.ReadLines("day10/input")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(10, 1, res)
	res = solvePart2(lines)
	utils.PrintSolution(10, 2, res)

}
