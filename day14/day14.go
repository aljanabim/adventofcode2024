package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

const (
	SIM_TIME = 100 // seconds
)

type Robot struct {
	X  int
	Y  int
	Vx int
	Vy int
}

type GridSize [2]int // X, Y

func (r *Robot) Step(n int, xLimit int, yLimit int) {
	r.X = (r.X + r.Vx*n) % xLimit
	if r.X < 0 { // loop the negative modulo
		r.X += xLimit
	}
	r.Y = (r.Y + r.Vy*n) % yLimit
	if r.Y < 0 { // loop the negative modulo
		r.Y += yLimit
	}
}

func parseNumSlice(numsStr []string) []int {
	nums := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		n, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
		if err != nil {
			panic(err)
		}
		nums[i] = int(n)
	}
	return nums
}

func parseLines(lines []string) []*Robot {
	robots := make([]*Robot, len(lines))
	for i, line := range lines {
		nums := parseNumSlice(strings.Split(strings.ReplaceAll(line[2:], " v=", ","), ","))
		robots[i] = &Robot{X: nums[0], Y: nums[1], Vx: nums[2], Vy: nums[3]}
	}
	return robots
}

// -1 for not in quadrant
// 0 for top left, 1 for top right
// 2 for bottom right, 3 for bottom right
func getQuadrant(robot *Robot, xLimit, yLimit int) int {
	leftQuadXUpperLim := (xLimit - 1) / 2 // exclusive
	topQuadYUpperLim := (yLimit - 1) / 2  // exclusive

	if robot.X < leftQuadXUpperLim && robot.Y < topQuadYUpperLim { // top left
		return 0
	} else if robot.X > leftQuadXUpperLim && robot.Y < topQuadYUpperLim { // top right
		return 1
	} else if robot.X < leftQuadXUpperLim && robot.Y > topQuadYUpperLim { // bottom left
		return 2
	} else if robot.X > leftQuadXUpperLim && robot.Y > topQuadYUpperLim { // bottom right
		return 3
	}

	return -1
}

func solvePart1(lines []string, sim_steps, xLimit, yLimit int) int {
	quadrants := [4]int{0, 0, 0, 0}
	robots := parseLines(lines)

	for _, robot := range robots {
		robot.Step(sim_steps, xLimit, yLimit)
		if q := getQuadrant(robot, xLimit, yLimit); q >= 0 {
			quadrants[q] += 1
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func solvePart2(lines []string) int {
	// sim_steps := 100
	xLimit := 101
	yLimit := 103
	grid := [103][101]string{}
	robots := parseLines(lines)

	// scanner := bufio.NewScanner(os.Stdin)

	i := 0
	for {
		i++
		xPosCount := map[int]int{}
		yPosCount := map[int]int{}
		for row := range yLimit {
			for col := range xLimit {
				grid[row][col] = "."
			}
		}
		for _, robot := range robots {
			robot.Step(1, xLimit, yLimit)
			grid[robot.Y][robot.X] = "X"
			xPosCount[robot.X]++
			yPosCount[robot.Y]++
		}
		fmt.Println("step", i)
		xPosOk := false
		yPosOk := false
		for _, count := range xPosCount {
			if count >= 31 {
				xPosOk = true
				break
			}
		}

		for _, count := range yPosCount {
			if count >= 31 {
				yPosOk = true
				break
			}
		}

		if xPosOk && yPosOk {
			for _, row := range grid {
				fmt.Println(row)
			}
			fmt.Println("step", i)
			reader := bufio.NewReader(os.Stdin)
			_, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	return i
}

func Solve() {
	lines, err := utils.ReadLines("day14/input.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines, 100, 101, 103)
	utils.PrintSolution(14, 1, res)
	res = solvePart2(lines)
	utils.PrintSolution(14, 2, res)
	// too low
	// 103401760
}
