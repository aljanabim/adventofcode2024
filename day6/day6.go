package day6

import (
	"fmt"

	"github.com/aljanabim/adventofcode2024/utils"
)

type obstacles map[[2]int]int

var directions = map[int][2]int{
	0: {-1, 0}, // up
	1: {0, 1},  // right
	2: {1, 0},  // down
	3: {0, -1}, // left
}

func step(pos [2]int, direction int) [2]int {
	newPos := [2]int{}
	newPos[0] = pos[0] + directions[direction][0]
	newPos[1] = pos[1] + directions[direction][1]
	return newPos

}

func solvePart1(obstacles obstacles, pos [2]int, rows, cols int) int {
	visits := map[[2]int]bool{}
	currentDirection := 0
	obscount := 0
	for inside(pos, rows, cols) {
		visits[pos] = true
		newPos := step(pos, currentDirection)
		if _, ok := obstacles[newPos]; ok {
			currentDirection = (currentDirection + 1) % len(directions)
			obscount++
			continue
		}
		pos = newPos

	}
	return len(visits)
}

func topCheck(obs1, obs2 [3]int) bool {
	return obs1[0]+1 == obs2[0] && obs1[1] < obs2[1]
}
func rightCheck(obs1, obs2 [3]int) bool {
	return obs1[0] < obs2[0] && obs1[1]-1 == obs2[1]
}
func leftCheck(obs1, obs2 [3]int) bool {
	return obs1[0] > obs2[0] && obs1[1]+1 == obs2[1]
}
func bottomCheck(obs1, obs2 [3]int) bool {
	return obs2[1] < obs1[1] && obs1[0]-1 == obs2[0]
}

func checkObstacleLoop(obs1, obs2, obs3 [3]int, visits map[[2]int]bool) bool {
	dir := obs1[2]
	if dir == 0 { // obs1 is top left
		tc := topCheck(obs1, obs2)
		rc := rightCheck(obs2, obs3)
		newObs := [2]int{obs3[0] - 1, obs1[1] - 1}
		_, onPath := visits[newObs]

		if tc && rc && onPath {
			fmt.Println("dir", dir, "obs1", obs1, "new obstacle at", newObs)
			return true
		}
	} else if dir == 1 { // obs1 is top right
		rc := rightCheck(obs1, obs2)
		bc := bottomCheck(obs2, obs3)
		newObs := [2]int{obs1[0] - 1, obs3[1] + 1}
		_, onPath := visits[newObs]
		if rc && bc && onPath {
			fmt.Println("dir", dir, "obs1", obs1, "new obstacle at", newObs)
			return true
		}

	} else if dir == 2 { // obs1 is bottom right
		tc := bottomCheck(obs1, obs2)
		rc := leftCheck(obs2, obs3)
		newObs := [2]int{obs3[0] + 1, obs1[1] + 1}
		_, onPath := visits[newObs]

		if tc && rc && onPath {
			fmt.Println("dir", dir, "obs1", obs1, "new obstacle at", newObs)
			return true
		}

	} else if dir == 3 { // obs1 is bottom left
		tc := leftCheck(obs1, obs2)
		rc := topCheck(obs2, obs3)
		newObs := [2]int{obs1[0] + 1, obs3[1] - 1}
		_, onPath := visits[newObs]

		if tc && rc && onPath {
			fmt.Println("dir", dir, "obs1", obs1, "new obstacle at", newObs)
			return true
		}
	}

	return false
}

/*
....#.....
....xxxxx#
....x...x.
..#.x...x.
..xxxxx#x.
..x.x.x.x.
.#xxxxxxx.
.xxxxxxx#.
#xxxxxxx..
......#x..

....#.....
....01111#
....0...2.
..#.0...2.
..01011#2.
..0.0.2.2.
.#3o^3332.
.01111oo#.
#o3o3322..
......#o..
*/

func isLoop(obstacles obstacles, newObstacle [2]int, startPos [2]int, rows, cols int) bool {
	MAXSTEPS := rows * cols
	steps := 0
	dir := 0
	pos := startPos

	for {
		if !inside(pos, rows, cols) {
			return false
		}
		newPos := step(pos, dir)
		if _, ok := obstacles[newPos]; ok || newPos == newObstacle {
			dir = (dir + 1) % len(directions)
			continue
		}
		pos = newPos
		steps++
		if steps > MAXSTEPS {
			return true
		}
	}
}

func solvePart2(obstacles obstacles, pos [2]int, rows, cols int) (int, map[[2]int]int) {
	visits := [][3]int{}
	startPos := pos
	currentDirection := 0

	newObstacles := map[[2]int]int{}
	newObsCount := 0
	for inside(pos, rows, cols) {
		visits = append(visits, [3]int{pos[0], pos[1], currentDirection})
		newPos := step(pos, currentDirection)
		isNewPosObs := false
		if _, ok := obstacles[newPos]; ok {
			isNewPosObs = true
			currentDirection = (currentDirection + 1) % len(directions)
		}
		if !isNewPosObs {
			pos = newPos
		}
	}
	for _, visit := range visits {
		pos := [2]int(visit[:2])
		dir := visit[2]
		newObs := step(pos, dir)
		newDir := (dir + 1) % len(directions)

		foundCompatableObs := false
		_, newObsExists := newObstacles[newObs]
		if _, ok := obstacles[newObs]; !ok && inside(newObs, rows, cols) && newObs != startPos && !newObsExists { // newObstacle cannot be existing obstacle
			for obs, _ := range obstacles {
				if newDir == 0 {
					if obs[0] < pos[0] && obs[1] == pos[1] {
						loopCheck := isLoop(obstacles, newObs, startPos, rows, cols)
						foundCompatableObs = loopCheck
					}
				} else if newDir == 1 {
					if obs[0] == pos[0] && obs[1] > pos[1] {
						loopCheck := isLoop(obstacles, newObs, startPos, rows, cols)
						foundCompatableObs = loopCheck
					}
				} else if newDir == 2 {
					if obs[0] > pos[0] && obs[1] == pos[1] {
						loopCheck := isLoop(obstacles, newObs, startPos, rows, cols)
						foundCompatableObs = loopCheck
					}
				} else if newDir == 3 {
					if obs[0] == pos[0] && obs[1] < pos[1] {
						loopCheck := isLoop(obstacles, newObs, startPos, rows, cols)
						foundCompatableObs = loopCheck
					}
				}
				if foundCompatableObs {
					newObstacles[newObs] += 1
					break
				}
			}
		}
		if foundCompatableObs {
			newObsCount++
		}
	}

	return newObsCount, newObstacles

}

func inside(pos [2]int, rows, cols int) bool {
	if pos[1] < 0 || pos[1] >= cols {
		return false
	}
	if pos[0] < 0 || pos[0] >= rows {
		return false
	}
	return true
}

func buildGrid(lines []string) (int, int, obstacles, [2]int) {
	gridObstacles := obstacles{}
	var currentPos [2]int
	for row, line := range lines {
		for col, cell := range line {
			if cell == '#' {
				gridObstacles[[2]int{row, col}] = -1
			} else if cell == '^' {
				currentPos = [2]int{row, col}
			}
		}
	}
	return len(lines), len(lines[0]), gridObstacles, currentPos
}

func Solve() {
	lines, err := utils.ReadLines("day6/input")
	if err != nil {
		panic(err)
	}

	rows, cols, gridObstacles, currPos := buildGrid(lines)

	res := solvePart1(gridObstacles, currPos, rows, cols)
	utils.PrintSolution(6, 1, res)
	res, obses := solvePart2(gridObstacles, currPos, rows, cols)
	for obs, count := range obses {
		if count > 1 {
			fmt.Println(obs)
		}
	}
	// for _, obs := range newObs {
	// 	lines[obs[0]] = lines[obs[0]][:obs[1]] + "O" + lines[obs[0]][obs[1]+1:]
	// }
	// for _, line := range lines {
	// 	fmt.Println(line)

	// }
	utils.PrintSolution(6, 2, res)
}

/*
....#.....
.........#
..........
..#.......
.......#..
..........
.#.O^.....
......OO#.
#O.O......
......#O..
*/
