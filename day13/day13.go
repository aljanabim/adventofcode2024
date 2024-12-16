package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Vec [2]float64

func (m *Vec) length() float64 {
	return math.Sqrt(m[0]*m[0] + m[1]*m[1])
}

type Matrix [2][2]float64

func (m *Matrix) dot(v Vec) Vec {
	return Vec{
		m[0][0]*v[0] + m[0][1]*v[1],
		m[1][0]*v[0] + m[1][1]*v[1],
	}
}

type Machine struct {
	Target Vec
	Motion Matrix
}

func inverse(mat Matrix) (Matrix, error) {
	inv := Matrix{}
	a := mat[0][0]
	b := mat[0][1]
	c := mat[1][0]
	d := mat[1][1]
	det := a*d - b*c
	if det == 0 {
		return inv, fmt.Errorf("determinant is 0, cannot compute inverse")
	}

	inv[0][0] = d / det
	inv[0][1] = -b / det
	inv[1][0] = -c / det
	inv[1][1] = a / det

	return inv, nil
}

func parseNumSlice(numsStr []string) []float64 {
	nums := make([]float64, len(numsStr))
	for i, numStr := range numsStr {
		n, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func parseLines(lines []string, part2 bool) []Machine {
	machines := []Machine{}
	for row, line := range lines {
		if len(line) == 0 || row == 0 {
			shift := 1
			if row == 0 {
				shift = 0
			}
			// Button A
			buttonA := parseNumSlice(strings.Split(strings.ReplaceAll(lines[row+shift][12:], "Y+", ""), ","))
			// Button B
			buttonB := parseNumSlice(strings.Split(strings.ReplaceAll(lines[row+1+shift][12:], "Y+", ""), ","))
			// Prize
			Prize := parseNumSlice(strings.Split(strings.ReplaceAll(lines[row+2+shift][9:], "Y=", ""), ","))

			target := Vec{Prize[0], Prize[1]}
			if part2 {
				target[0] += 10000000000000
				target[1] += 10000000000000
			}
			machines = append(machines,
				Machine{
					Target: target,
					Motion: Matrix{
						{buttonA[0], buttonB[0]},
						{buttonA[1], buttonB[1]}},
				})
		}
	}
	return machines
}

func round(n float64, d int) float64 {
	dec := math.Pow(10, float64(d))
	return math.Round(n*dec) / dec
}

func solveDay(lines []string, part2 bool) int {
	defer utils.Duration(utils.Track(fmt.Sprintf("Solve Day, is Part 2 %v", part2)))
	machines := parseLines(lines, part2)
	cost := 0
	for _, m := range machines {
		inv, err := inverse(m.Motion)
		if err != nil {
			continue
		}
		sol := inv.dot(m.Target)
		solTarget := m.Motion.dot(Vec{math.Round(sol[0]), math.Round(sol[1])})
		distVec := Vec{m.Target[0] - solTarget[0], m.Target[1] - solTarget[1]}
		solValid := distVec.length() == 0 && sol[0] >= 0 && sol[1] >= 0
		if solValid {
			buttonAPresses := sol[0]
			buttonBPresses := sol[1]
			cost += int(math.Round(buttonAPresses)*3) + int(math.Round(buttonBPresses))
		}
	}
	return cost
}

func Solve() {
	lines, err := utils.ReadLines("day13/input")
	if err != nil {
		panic(err)
	}

	res := solveDay(lines, false)
	utils.PrintSolution(13, 1, res)
	res = solveDay(lines, true)
	utils.PrintSolution(13, 2, res)
}
