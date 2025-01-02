package day23

import (
	"slices"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Graph struct {
	nodes []string
	edges [][2]string
}

func (g *Graph) addNode(name string) {
	nodeExists := false
	for _, n := range g.nodes {
		if n == name {
			nodeExists = true
		}
	}
	if !nodeExists {
		g.nodes = append(g.nodes, name)
	}
}

func (g *Graph) addEdge(from, to string) {
	edgeExists := false
	for _, edge := range g.edges {
		if edge[0] == from && edge[1] == to || edge[0] == to && edge[1] == from {
			edgeExists = true
		}
	}
	if !edgeExists {
		g.edges = append(g.edges, [2]string{from, to})
	}

	// Add any missing node
	fromNodeExists := false
	toNodeExists := false
	for _, node := range g.nodes {
		if node == from {
			fromNodeExists = true
		}
		if node == to {
			toNodeExists = true
		}
	}
	if !fromNodeExists {
		g.addNode(from)
	}
	if !toNodeExists {
		g.addNode(to)
	}
}

func (g *Graph) getNeighbors(node string) []string {
	n := []string{}
	for _, edge := range g.edges {
		if edge[0] == node {
			n = append(n, edge[1])
		} else if edge[1] == node {
			n = append(n, edge[0])
		}
	}
	return n
}

// func (g *Graph) findCliques(node, prevNode string, size int, startNode string, visited map[[2]string]bool) int {
// 	visited[[2]string{prevNode, node}] = true
// 	visited[[2]string{node, prevNode}] = true

// 	if size == 0 {
// 		fmt.Println("Bummer")
// 		return 0
// 	}
// 	neighbours := g.getNeighbors(node)
// 	fmt.Println("currentNode", node, "neigbours", neighbours, "level", size)
// 	cliques := 0
// 	for _, n := range neighbours {
// 		if n == startNode && size == 1 {
// 			fmt.Println("Found startNode from", node, size)
// 			cliques++
// 			continue
// 		}
// 		if !visited[[2]string{node, n}] {
// 			fmt.Printf("Source %s has neighbor %s at level %d\n", node, n, size)
// 			cliques += g.findCliques(n, node, size-1, startNode, visited)
// 		}
// 	}
// 	return cliques
// }

type Record struct {
	string
	int
}

// func (g *Graph) findCliques(startNode string, size int) int {
// 	cliques := 0

// 	deque := []Record{{startNode, 0}}
// 	// visited := map[string]bool{}
// 	for len(deque) > 0 {
// 		node := deque[0]
// 		// visited[node.string] = true
// 		deque = deque[1:]
// 		if node.int < size {
// 			neighbors := g.getNeighbors(node.string)
// 			for _, n := range neighbors {
// 				// if !visited[n] {
// 				fmt.Println("From", node, "visit", n)
// 				deque = append(deque, Record{n, node.int + 1})
// 				// visited[n] = true
// 				// }
// 			}
// 		}
// 	}

//		return cliques
//	}
func reverse[T any](values []T) []T {
	newSlice := make([]T, len(values))
	for i, j := 0, len(values)-1; i <= j; i, j = i+1, j-1 {
		newSlice[i], newSlice[j] = values[j], values[i]
	}
	return newSlice
}

// func flip(s string) string {
// 	rs := []byte(s)
// 	slices.Reverse(rs)
// 	return string(rs)
// }

func (g *Graph) findCliques(node string, size int, visited map[string]bool) int {
	// if node == startNode && size == 0 {
	// 	if results[[3]byte(sequence)] || results[[3]byte(reverse(sequence))] {
	// 		return 0
	// 	}
	// 	results[[3]byte(sequence)] = true
	// 	return 1
	// }

	// loop := [][]rune{}
	// neighbors := g.getNeighbors(node)
	// fmt.Println("Node", node)
	// fmt.Println("Visiting neighbor", 2)
	// path := []string{node, "2"}
	cycles := [][]string{}
	visited[node] = true
	for _, n := range g.getNeighbors(node) {
		if !visited[n] {
			g.dfs(n, node, size-1, []string{node, n}, &cycles, visited)
		}
	}
	// for _, c := range cycles {
	// 	fmt.Println(c)
	// }
	// for _, n := range neighbors {
	// }

	return len(cycles)
}

func (g *Graph) dfs(node string, start string, depth int, path []string, cycles *[][]string, visited map[string]bool) bool {
	if depth == 0 {
		return false
	}
	for _, n := range g.getNeighbors(node) {
		if n == start && depth == 1 {
			path = append(path, n)
			// fmt.Println("== found start", n, path)
			add := true
			joined := strings.Join(path, "")
			joinedRev := strings.Join(reverse(path), "")

			for _, c := range *cycles {
				cycleJoined := strings.Join(c, "")
				if cycleJoined == joined || cycleJoined == joinedRev {
					add = false
				}
			}

			if add {
				newPath := make([]string, len(path))
				copy(newPath, path)
				*cycles = append(*cycles, newPath)
			}
			return true
		}
		if !slices.Contains(path, n) && !visited[n] {
			// fmt.Println(">> from", node, "visit", n, path)
			path = append(path, n)
			g.dfs(n, start, depth-1, path, cycles, visited)
			// back track path
			path = path[:len(path)-1]

		}
	}
	return false
}

func buildGraph(lines []string) *Graph {
	graph := Graph{}
	for _, line := range lines {
		edge := strings.Split(line, "-")
		graph.addEdge(edge[0], edge[1])
	}
	return &graph
}

func solvePart1(lines []string) int {
	graph := buildGraph(lines)
	defer utils.Duration(utils.Track("Part 1"))
	visited := map[string]bool{}
	count := 0
	for _, node := range graph.nodes {
		if node[0] == 't' {
			// fmt.Println(node, "->", graph.getNeighbors(node))
			cliques := graph.findCliques(node, 3, visited)
			count += cliques
			// fmt.Printf("Found %d cliques for %s\n", len(cliques), node)
			// for _, c := range cliques {
			// 	fmt.Println(c)
			// }
		}
	}
	return count
}

func Solve() {
	lines, err := utils.ReadLines("day23/input.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(23, 1, res)

}
