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

type Record struct {
	string
	int
}

func reverse[T any](values []T) []T {
	newSlice := make([]T, len(values))
	for i, j := 0, len(values)-1; i <= j; i, j = i+1, j-1 {
		newSlice[i], newSlice[j] = values[j], values[i]
	}
	return newSlice
}

func (g *Graph) findCliques(node string, size int, visited map[string]bool) int {
	cycles := [][]string{}
	visited[node] = true
	for _, n := range g.getNeighbors(node) {
		if !visited[n] {
			g.dfs(n, node, size-1, []string{node, n}, &cycles, visited)
		}
	}
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

func (g *Graph) findLargestClique(node string, size int, visited map[string]bool) int {
	cycles := [][]string{}
	visited[node] = true
	for _, n := range g.getNeighbors(node) {
		if !visited[n] {
			g.dfsMax(n, node, size-1, []string{node, n}, &cycles, visited)
		}
	}
	return len(cycles)
}

func (g *Graph) dfsMax(node string, start string, depth int, path []string, cycles *[][]string, visited map[string]bool) bool {
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
			g.dfsMax(n, start, depth-1, path, cycles, visited)
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

func solvePart1(graph *Graph) int {
	defer utils.Duration(utils.Track("Part 1"))
	visited := map[string]bool{}
	count := 0
	for _, node := range graph.nodes {
		if node[0] == 't' {
			cliques := graph.findCliques(node, 3, visited)
			count += cliques
		}
	}
	return count
}

func Solve() {
	lines, err := utils.ReadLines("day23/input.txt")
	if err != nil {
		panic(err)
	}
	graph := buildGraph(lines)
	res := solvePart1(graph)
	utils.PrintSolution(23, 1, res)

}
