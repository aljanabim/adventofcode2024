package day18

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

var directions = []Direction{UP, RIGHT, DOWN, LEFT}

var dir2Vec = map[Direction][2]int{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

type Entity int

const (
	FREE Entity = iota
	WALL
)

func CreateGrid(lines []string, size int) ([][]Entity, error) {
	grid := make([][]Entity, size)
	for i := range size {
		row := make([]Entity, size)
		for j := range size {
			row[j] = FREE
		}
		grid[i] = row
	}

	for _, line := range lines {
		numsStr := strings.Split(line, ",")
		x, err := strconv.ParseInt(strings.TrimSpace(numsStr[0]), 10, 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseInt(strings.TrimSpace(numsStr[1]), 10, 64)
		if err != nil {
			return nil, err
		}
		grid[int(y)][int(x)] = WALL
	}
	return grid, nil
}

func isPointSafe(grid [][]Entity, point [2]int) bool {
	if point[0] < 0 || point[0] > len(grid)-1 {
		return false
	}
	if point[1] < 0 || point[1] > len(grid[0])-1 {
		return false
	}
	if grid[point[0]][point[1]] != FREE {
		return false
	}
	return true
}

func Argmin(list []int, value int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

func SearchPathDFSBad(grid [][]Entity, point, target [2]int, visited map[[2]int]bool) int {
	if point[0] == target[0] && point[1] == target[1] {
		return 0
	}
	visited[point] = true
	visitCosts := []int{}
	visitedPoints := [][2]int{}
	for _, dir := range directions {
		nextPoint := [2]int{point[0] + dir2Vec[dir][0], point[1] + dir2Vec[dir][1]}
		if !visited[nextPoint] && isPointSafe(grid, nextPoint) {
			visitCosts = append(visitCosts, 1+SearchPathDFSBad(grid, nextPoint, target, visited))
			visitedPoints = append(visitedPoints, nextPoint)
		}
	}
	if len(visitCosts) > 0 {
		minCost := slices.Min(visitCosts)
		fmt.Println("Good node", visitCosts, visitedPoints[Argmin(visitCosts, minCost)])
		return minCost
	}
	return 2 * len(grid)
}

func FindOptimalPath(grid [][]Entity, start, target [2]int) int {
	visited := map[[2]int]bool{start: true}
	queue := [][3]int{{start[0], start[1], 0}}
	for len(queue) > 0 {
		q := queue[0]
		cost := q[2]
		point := [2]int(q[:2])
		// fmt.Println("Exploring", point, "at depth", q[2])
		queue = queue[1:]
		if point == target {
			return cost
		}
		for _, dir := range directions {
			nextPoint := [2]int{point[0] + dir2Vec[dir][0], point[1] + dir2Vec[dir][1]}
			if !visited[nextPoint] && isPointSafe(grid, nextPoint) {
				visited[nextPoint] = true
				queue = append(queue, [3]int{nextPoint[0], nextPoint[1], cost + 1})
			}
		}
	}
	return -1
}

func printGrid(grid [][]Entity) string {
	s := strings.Builder{}
	for _, row := range grid {
		for _, e := range row {
			switch e {
			case FREE:
				s.WriteRune('.')
			case WALL:
				s.WriteRune('#')
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

func solvePart1(lines []string, size int) int {
	defer utils.Duration(utils.Track("Testing Solve Part 1"))
	grid, err := CreateGrid(lines, size)
	if err != nil {
		panic(err)
	}
	// fmt.Print(printGrid(grid))
	return FindOptimalPath(grid, [2]int{0, 0}, [2]int{size - 1, size - 1})
}

func solvePart2(lines []string, size int, idx int) string {
	defer utils.Duration(utils.Track("Testing Solve Part 1"))
	ret := 0
	for ret != -1 {
		idx++
		grid, err := CreateGrid(lines[:idx], size)
		if err != nil {
			panic(err)
		}
		ret = FindOptimalPath(grid, [2]int{0, 0}, [2]int{size - 1, size - 1})
	}
	return lines[idx-1]
}

func Solve() {
	lines, err := utils.ReadLines("day18/input.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines[:1024], 71)
	utils.PrintSolution(18, 1, res)
	res2 := solvePart2(lines, 71, 1024)
	utils.PrintSolution(18, 2, res2)

}
