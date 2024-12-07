package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func parseRule(rule string) (int, int, error) {
	nums := strings.Split(rule, "|")
	left, err := strconv.ParseInt(nums[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse first part of rule %s", rule)
	}
	right, err := strconv.ParseInt(nums[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse second part of rule %s", rule)
	}
	return int(left), int(right), nil

}

func solveBothParts(pageRules, updatedPages string) (int, int) {
	totMids := 0
	totMidsCorrected := 0
	printPages := strings.Split(updatedPages, "\n")
	for _, printPage := range printPages {
		rules := map[int]int{}
		printNumsStr := strings.Split(printPage, ",")
		printNums := make([]int, len(printNumsStr))
		for i, numStr := range printNumsStr {
			numi64, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				panic(err)
			}
			num := int(numi64)
			rules[num] = 0
			printNums[i] = num
		}

		rulesLines := strings.Split(pageRules, "\n")
		for pos := range len(rules) - 1 {
			for _, rule := range rulesLines {
				left, right, err := parseRule(rule)
				if err != nil {
					panic(err)
				}
				_, okLeft := rules[left]
				_, okRight := rules[right]
				pLeft := rules[left]
				if okLeft && okRight && pLeft == pos {
					rules[right] = pos + 1
				}
			}

		}
		// check the print page order
		correctOrder := true
		for i, pageNum := range printNums {
			if rules[pageNum] != i {
				correctOrder = false
			}
		}
		if correctOrder {
			midNum := printNums[len(printNums)/2]
			totMids += int(midNum)
		} else {
			sortedPages := make([]int, len(rules))
			for k, v := range rules {
				sortedPages[v] = k
			}
			midNum := sortedPages[len(printNums)/2]
			totMidsCorrected += int(midNum)
		}
	}
	return totMids, totMidsCorrected
}

func Solve() {
	fileRaw, err := os.ReadFile("day5/input")
	if err != nil {
		panic(err)
	}
	fileStr := string(fileRaw)
	sections := strings.Split(fileStr, "\n\n")
	pageRules := sections[0]
	updatedPages := sections[1]

	res, resCorrected := solveBothParts(pageRules, updatedPages)
	utils.PrintSolution(5, 1, res)
	utils.PrintSolution(5, 2, resCorrected)
}
