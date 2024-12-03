package day1

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func getSortedLists() ([]int, []int) {
	dataRaw, err := os.ReadFile("day1/input")
	dataStr := strings.Split(string(dataRaw), "\n")
	if err != nil {
		panic(err)
	}
	l1 := make([]int, len(dataStr))
	l2 := make([]int, len(dataStr))
	for i, line := range dataStr {
		nums := strings.Split(line, " ")
		n1, err := strconv.ParseInt(nums[0], 10, 32)
		if err != nil {
			panic(err)
		}
		n2, err := strconv.ParseInt(nums[len(nums)-1], 10, 32)
		if err != nil {
			panic(err)
		}

		l1[i] = int(n1)
		l2[i] = int(n2)
	}
	slices.Sort(l1)
	slices.Sort(l2)
	return l1, l2
}

func solvePart1(l1, l2 []int) int {
	diff := 0
	for i := range len(l1) {
		diff += int(math.Abs(float64(l1[i] - l2[i])))
	}
	return diff
}
func solvePart2(l1, l2 []int) int {
	simMap := make(map[int]int)
	for _, v := range l2 {
		simMap[v] += 1
	}
	var totSim int = 0
	for _, v := range l1 {
		if val, ok := simMap[v]; ok {
			totSim += v * val
		}
	}

	return totSim
}

func Solve() {
	l1, l2 := getSortedLists()
	res := solvePart1(l1, l2)
	utils.PrintSolution(1, 1, res)
	res = solvePart2(l1, l2)
	utils.PrintSolution(1, 2, res)
}
