package day3

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func solvePart1() int64 {
	dataRaw, err := os.ReadFile("day3/input")
	if err != nil {
		panic(err)
	}
	dataStr := string(dataRaw)
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	multis := r.FindAllString(dataStr, -1)
	var res int64 = 0
	for _, multi := range multis {
		numsStr := strings.Split(multi[4:len(multi)-1], ",")
		num1, err := strconv.ParseInt(numsStr[0], 10, 64)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.ParseInt(numsStr[1], 10, 64)
		if err != nil {
			panic(err)
		}
		res += num1 * num2
	}
	return res
}

func Solve() {
	res := solvePart1()
	utils.PrintSolution(3, 1, res)

}
