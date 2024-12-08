package day6

import (
	"fmt"

	"github.com/aljanabim/adventofcode2024/utils"
)

type obstacles map[[2]int]bool

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

// func turn()

func solvePart1(obstacles obstacles, pos [2]int, rows, cols int) int {
	visits := map[[2]int]bool{}
	currentDirection := 0
	obscount := 0
	for inside(pos, rows, cols) {
		visits[pos] = true
		newPos := step(pos, currentDirection)
		if obs, ok := obstacles[newPos]; ok {
			if obs {
				currentDirection = (currentDirection + 1) % len(directions)
				obscount++
				continue
			}
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

func isLoop(obstacles obstacles, pos [2]int, rows, cols int) bool {
	MAXSTEPS := rows * cols
	steps := 0
	dir := 0

	for {
		if !inside(pos, rows, cols) {
			return false
		}
		if steps > MAXSTEPS {
			return true
		}

		newPos := step(pos, dir)
		isNewPosObs := false
		if obs, ok := obstacles[newPos]; ok {
			if obs {
				isNewPosObs = true
				dir = (dir + 1) % len(directions)
			}
		}
		if !isNewPosObs {
			pos = newPos
		}
	}
}

func solvePart2(obstacles obstacles, pos [2]int, rows, cols int) (int, [][2]int) {
	visits := [][3]int{}
	// visits := map[[2]int]bool{}
	// obstaclesInPath := [][3]int{}
	obstaclesInPath := map[[2]int]int{}
	currentDirection := 0

	// newObsCount := 0
	newObstacles := [][2]int{}
	// obscount := 0
	newObsCount := 0
	for inside(pos, rows, cols) {
		// visits[pos] = true
		// only log visits after the first obstacle
		// if obscount > 0 {
		visits = append(visits, [3]int{pos[0], pos[1], currentDirection})
		// }
		newPos := step(pos, currentDirection)
		isNewPosObs := false
		if obs, ok := obstacles[newPos]; ok {
			if obs {
				isNewPosObs = true
				obstaclesInPath[newPos] = currentDirection
				currentDirection = (currentDirection + 1) % len(directions)
			}
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
		// fmt.Println(pos, "smash dir", dir, "new obs", newObs)
		if _, ok := obstaclesInPath[newObs]; !ok { // newObstacle cannot be existing obstacle
			for obs, smashDir := range obstaclesInPath {
				if newDir == 0 {
					if obs[0] < pos[0] && obs[1] == pos[1] && smashDir == newDir {
						foundCompatableObs = true
					}
				} else if newDir == 1 {
					if obs[0] == pos[0] && obs[1] > pos[1] && smashDir == newDir {
						foundCompatableObs = true
					}
				} else if newDir == 2 {
					if obs[0] > pos[0] && obs[1] == pos[1] && smashDir == newDir {
						foundCompatableObs = true
					}
				} else if newDir == 3 {
					if obs[0] == pos[0] && obs[1] < pos[1] && smashDir == newDir {
						foundCompatableObs = true
					}
				}
				if newObs == [2]int{4, 8} {
					fmt.Println(pos, "smash dir", dir, "new obs", newObs, "compatible obs", obs)
				}
				if foundCompatableObs {
					fmt.Println(">> smash dir", dir, "new obs", newObs, "compatible obs", obs)
					newObstacles = append(newObstacles, newObs)
					break
				}
			}
			// fmt.Println("suggested obs", newObs, "is new")
			// fmt.Println("looking for obstacle with", "row", newObs[0]+1, "col >", newObs[1])
			// fmt.Println("pos", pos, "dir", dir, newDir, newObs, foundCompatableObs)
		}
		if foundCompatableObs {
			newObsCount++
		}
	}
	fmt.Println("New obs count", newObsCount)

	// fmt.Println(obstaclesInPath)

	// loopCount := 0
	// for i := range len(obstaclesInPath) - 2 {
	// 	if checkObstacleLoop(obstaclesInPath[i], obstaclesInPath[i+1], obstaclesInPath[i+2], visits) {
	// 		loopCount += 1
	// 	}
	// }
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
				gridObstacles[[2]int{row, col}] = true
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
	res, _ = solvePart2(gridObstacles, currPos, rows, cols)
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
