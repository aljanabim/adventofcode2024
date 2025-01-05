package day16

import (
	"fmt"
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
	Id      int
	Heading Heading
	Pos     [2]int
	Up      *Node
	Right   *Node
	Down    *Node
	Left    *Node
}

type Heading int

const (
	UP Heading = iota
	RIGHT
	DOWN
	LEFT
	MAX_HEADING
)

func createMaze(lines []string) Maze {
	maze := Maze{Height: len(lines) - 1, Width: len(lines[0]), Nodes: map[[2]int]*Node{}}
	// Create nodes in maze
	id := 0
	for row, line := range lines {
		for col, cell := range line {
			node := &Node{Pos: [2]int{row, col}, Id: id, Heading: RIGHT}
			if cell != '#' {
				maze.Nodes[[2]int{row, col}] = node
				id++
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
func dfs(node *Node, depth, cost int, path []*Node, dist []int, heading Heading, visited map[[2]int]bool) {
	fmt.Println("At", node.Pos, "headed", heading, "cost", cost, "depth", depth)
	visited[node.Pos] = true
	if depth > 0 && slices.Contains(path, node) {
		fmt.Println("Reached a pre visited point", node.Pos, cost, "opt cost", dist[node.Id])
		return
	}

	// if node == endNode {
	// 	fmt.Println("Reached end", cost, *minCost, len(path))
	// 	if cost < *minCost {
	// 		*minCost = cost
	// 	}
	// 	return
	// }

	if node.Up != nil && !visited[node.Up.Pos] {
		turnCost := 0
		if heading != UP {
			turnCost = 1000
		}
		newCost := cost + 1 + turnCost
		dfs(node.Up, depth+1, newCost, path, dist, UP, visited)
	}
	if node.Right != nil && !visited[node.Right.Pos] {
		turnCost := 0
		if heading != RIGHT {
			turnCost = 1000
		}
		newCost := cost + 1 + turnCost
		dfs(node.Right, depth+1, newCost, path, dist, RIGHT, visited)
	}
	if node.Down != nil && !visited[node.Down.Pos] {
		turnCost := 0
		if heading != DOWN {
			turnCost = 1000
		}
		newCost := cost + 1 + turnCost
		dfs(node.Down, depth+1, newCost, path, dist, DOWN, visited)
	}
	if node.Left != nil && !visited[node.Left.Pos] {
		turnCost := 0
		if heading != LEFT {
			turnCost = 1000
		}
		newCost := cost + 1 + turnCost
		dfs(node.Left, depth+1, newCost, path, dist, LEFT, visited)
	}
}

func computeCost(source *Node, target *Node, heading Heading, queue, dist, prev []int) {
	if target != nil && slices.Contains(queue, target.Id) {
		target.Heading = heading
		alt := dist[source.Id] + 1
		if source.Heading != target.Heading {
			alt += 1000
		}
		if alt < dist[target.Id] {
			dist[target.Id] = alt
			prev[target.Id] = source.Id
		}
		// fmt.Println("At", source.Pos, dist[source.Id], "going", heading, target.Pos, dist[target.Id])
	}

}

func dijkstra(nodes []*Node, source, target *Node, inf int) ([]int, []int) {
	dist := make([]int, len(nodes))
	prev := make([]int, len(nodes))
	queue := []int{}
	for _, node := range nodes {
		dist[node.Id] = inf
		prev[node.Id] = inf
		queue = append(queue, node.Id)
	}
	dist[source.Id] = 0
	for len(queue) > 0 {
		// get vertex in queue with minimum dist
		var n *Node
		minCost := inf
		for _, id := range queue {
			if dist[id] < minCost {
				minCost = dist[id]
				n = nodes[id]
			}
		}
		// Correction shifting costs due to turning one step back
		if prev[n.Id] != inf {
			if n.Heading != nodes[prev[n.Id]].Heading {
				dist[prev[n.Id]] += 1000
			}
		}
		if n == target {
			break
		}
		nIdx := slices.Index(queue, n.Id)
		queue = slices.Delete(queue, nIdx, nIdx+1)
		// for each neighbor v of u still in Q:
		computeCost(n, n.Up, UP, queue, dist, prev)
		computeCost(n, n.Right, RIGHT, queue, dist, prev)
		computeCost(n, n.Down, DOWN, queue, dist, prev)
		computeCost(n, n.Left, LEFT, queue, dist, prev)
	}

	return dist, prev
}
func solvePart1(lines []string) int {
	defer utils.Duration(utils.Track("Part 1"))
	maze := createMaze(lines)
	nodes := make([]*Node, len(maze.Nodes))
	for _, node := range maze.Nodes {
		nodes[node.Id] = node
	}

	dist, _ := dijkstra(nodes, maze.Start, maze.End, maze.Height*maze.Width*1000)
	return dist[maze.End.Id]
}

func solvePart2(lines []string) int {
	defer utils.Duration(utils.Track("Part 2"))
	maze := createMaze(lines)
	nodes := make([]*Node, len(maze.Nodes))
	for _, node := range maze.Nodes {
		nodes[node.Id] = node
	}

	dist, prev := dijkstra(nodes, maze.Start, maze.End, maze.Height*maze.Width*1000)
	// distToEnd := dist[maze.End.Id]
	// slices.Sort(dist)
	nextId := maze.End.Id
	optimalPath := []*Node{}
	for nextId != maze.Start.Id {
		optimalPath = append(optimalPath, nodes[nextId])
		nextId = prev[nextId]
	}
	optimalPath = append(optimalPath, maze.Start)
	slices.Reverse(optimalPath)
	// Find all optimal paths
	visited := map[[2]int]bool{optimalPath[0].Pos: true}
	for i, step := range optimalPath[1:3] {
		visited[step.Pos] = true
		fmt.Println("Starting from", step.Pos, optimalPath[i].Heading)
		dfs(step, 0, dist[step.Id], optimalPath, dist, optimalPath[i].Heading /*heading of prev step*/, visited)
		fmt.Println(step.Pos, dist[step.Id])
	}
	return dist[maze.End.Id]
}

func Solve() {
	lines, err := utils.ReadLines("day16/input_mini3.txt")
	if err != nil {
		panic(err)
	}
	// res := solvePart1(lines)
	// utils.PrintSolution(16, 1, res)
	res := solvePart2(lines)
	utils.PrintSolution(16, 2, res)
}
