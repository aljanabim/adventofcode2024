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

func traverse(n *Node) bool {
	n.visited = true
	if n.height == 0 {
		return true
	}
	if n.up != nil && !n.up.visited {
		return traverse(n.up)
	}
	if n.down != nil && !n.down.visited {
		return traverse(n.down)
	}
	if n.left != nil && !n.left.visited {
		return traverse(n.left)
	}
	if n.right != nil && !n.right.visited {
		return traverse(n.right)
	}
	return false
}

func resetVisited(n *Node) {
	if n.visited {
		n.visited = false
	}
	if n.up != nil && n.up.visited {
		resetVisited(n.up)
	}
	if n.down != nil && n.down.visited {
		resetVisited(n.down)
	}
	if n.left != nil && n.left.visited {
		resetVisited(n.left)
	}
	if n.right != nil && n.right.visited {
		resetVisited(n.right)
	}

}

func solvePart1(lines []string) int {
	grid := make([][]*Node, len(lines))
	peakNodes := []*Node{}

	for row, line := range lines {
		r := []*Node{}
		for _, c := range line {
			r = append(r, &Node{height: int(c - '0')})
		}
		grid[row] = r
	}
	// Populate grid
	for i, row := range grid {
		for j, node := range row {
			if j > 0 { // left neighbor
				leftNode := grid[i][j-1]
				if utils.Abs(leftNode.height-node.height) == 1 {
					node.left = leftNode
					leftNode.right = node
				}
			}
			if j < len(row)-1 { // right neighbor
				rightNode := grid[i][j+1]
				if utils.Abs(rightNode.height-node.height) == 1 {
					node.right = rightNode
					rightNode.left = node
				}
			}

			if i > 0 { // up neighbor
				upNode := grid[i-1][j]
				if utils.Abs(upNode.height-node.height) == 1 {
					node.up = upNode
					upNode.down = node
				}
			}
			if i < len(grid)-1 { // down neighbor
				downNode := grid[i+1][j]
				if utils.Abs(downNode.height-node.height) == 1 {
					node.down = downNode
					downNode.up = node
				}
			}

			if node.height == 9 {
				peakNodes = append(peakNodes, node)
			}
		}
	}
	/*
	   ..90..9
	   ...1.98
	   ...2..7
	   6543456
	   765.987
	   876....
	   987....
	*/

	score := 0
	for _, n := range peakNodes {
		reachable := traverse(n)
		resetVisited(n)
		if reachable {
			score++
		}
	}
	return score
}
