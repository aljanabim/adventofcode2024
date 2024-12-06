package day5

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func checkLines(rules map[int64]int, line string) (bool, int) {
	nums := strings.Split(line, ",")
	for i := range len(nums) - 1 {
		num1, err := strconv.ParseInt(nums[i], 10, 64)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.ParseInt(nums[i+1], 10, 64)
		if err != nil {
			panic(err)
		}
		if _, ok := rules[num1]; !ok {
			panic(fmt.Sprintf("No num1 one found %d", num1))
		}
		if _, ok := rules[num2]; !ok {
			panic(fmt.Sprintf("No num2 one found %d", num2))
		}
		if rules[num2] < rules[num1] {
			return false, -1
		}
	}
	midNumStr := nums[len(nums)/2]
	midNum, err := strconv.ParseInt(midNumStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return true, int(midNum)
}

func solvePart1(pageRules, updatedPages string) int {
	rules := map[int64]int{}
	lines := strings.Split(pageRules, "\n")
	for _, rule := range lines {
		nums := strings.Split(rule, "|")
		left, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			panic(err)
		}
		right, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			panic(err)
		}
		_, foundLeft := rules[left]
		_, foundRight := rules[right]

		if !foundLeft && !foundRight {
			rules[left] = 0
			rules[right] = 1
		} else if foundLeft && !foundRight {
			rules[right] = rules[left] + 1
		} else if foundLeft && foundRight {
			if rules[right] <= rules[left] {
				diff := int(math.Abs(float64(rules[right] - rules[left])))
				for num, pos := range rules {
					if num != left && num != right && pos >= rules[right] {
						rules[num] = pos + diff + 1
					}
				}
				rules[right] = rules[right] + 1 + diff
			}
		}
	}
	out := make([]int64, len(rules))
	for k, v := range rules {
		rules[k] = v % len(rules)
		fmt.Println(rules[k], k)
		out[rules[k]] = k
	}
	fmt.Println(out)

	numLines := strings.Split(updatedPages, "\n")
	sumMids := 0
	for _, line := range numLines {
		ordered, midNum := checkLines(rules, line)
		if ordered {
			sumMids += midNum
		}
	}
	return sumMids
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

	pageRules = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`

	updatedPages = `75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	res := solvePart1(pageRules, updatedPages)
	utils.PrintSolution(5, 1, res)
}
