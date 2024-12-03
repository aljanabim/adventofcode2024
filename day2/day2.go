package day2

import (
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

// checkReport returns true if the report
// is safe based on the following criteria
//  1. The levels are either all increasing
//     or all decreasing.
//  2. Any two adjacent levels differ by at
//     least one and at most three
func checkReport(report []int) bool {
	isAccending := report[0] < report[1]
	for i := range len(report) - 1 {
		if report[i] < report[i+1] != isAccending {
			return false
		}
		diff := math.Abs(float64(report[i] - report[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
func checkReportWithDampener(report []int) bool {
	firstCheck := checkReport(report)
	if firstCheck {
		return true
	} else {
		for n := range len(report) {
			shortReport := append([]int{}, report[:n]...)
			shortReport = append(shortReport, report[n+1:]...)
			if checkReport(shortReport) {
				return true
			}
		}
	}
	return false
}

func readReports() [][]int {
	dataRaw, err := os.ReadFile("day2/input")
	if err != nil {
		panic(err)
	}
	reportsStr := strings.Split(string(dataRaw), "\n")
	reports := make([][]int, len(reportsStr))
	for reportIdx, reportStr := range reportsStr {
		levelsStr := strings.Split(reportStr, " ")
		levels := make([]int, len(levelsStr))
		for levelIdx, levelStr := range levelsStr {
			level, err := strconv.ParseInt(levelStr, 10, 64)
			if err != nil {
				panic(err)
			}
			levels[levelIdx] = int(level)
		}
		reports[reportIdx] = levels
	}

	return reports
}

func solvePart1(reports [][]int) int {
	totSafe := 0
	for _, report := range reports {
		if checkReport(report) {
			totSafe += 1
		}
	}
	return totSafe
}
func solvePart2(reports [][]int) int {
	totSafe := 0
	for _, report := range reports {
		if checkReportWithDampener(report) {
			totSafe += 1
		}
	}
	return totSafe
}
func Solve() {
	reports := readReports()
	res := solvePart1(reports)
	utils.PrintSolution(2, 1, res)
	res = solvePart2(reports)
	utils.PrintSolution(2, 2, res)

}
