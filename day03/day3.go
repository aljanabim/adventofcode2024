package day03

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func parseMulti(multi string) int64 {
	numsStr := strings.Split(multi[4:len(multi)-1], ",")
	num1, err := strconv.ParseInt(numsStr[0], 10, 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseInt(numsStr[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return num1 * num2
}

func solvePart1() int64 {
	dataRaw, err := os.ReadFile("day03/input")
	if err != nil {
		panic(err)
	}
	dataStr := string(dataRaw)
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	multis := r.FindAllString(dataStr, -1)
	var res int64 = 0
	for _, multi := range multis {
		res += parseMulti(multi)
	}
	return res
}

func solvePart2() int64 {
	dataRaw, err := os.ReadFile("day03/input")
	if err != nil {
		panic(err)
	}
	dataStr := string(dataRaw)
	// r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	r := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\(\d+,\d+\)`)

	multis := r.FindAllString(dataStr, -1)
	var res int64 = 0
	do := true
	for _, multi := range multis {
		if strings.Contains(multi, "'") {
			do = false
		} else if strings.Contains(multi, "do(") {
			do = true
		} else if strings.Contains(multi, "mul") && do {

			res += parseMulti(multi)
		}
	}
	return res
}

func Solve() {
	res := solvePart1()
	utils.PrintSolution(3, 1, res)
	res = solvePart2()
	utils.PrintSolution(3, 2, res)

}
