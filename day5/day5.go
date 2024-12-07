package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func checkLines(rules map[int]int, line string) (bool, int) {
	nums := strings.Split(line, ",")
	for i := range len(nums) - 1 {
		num1i64, err := strconv.ParseInt(nums[i], 10, 64)
		if err != nil {
			panic(err)
		}
		num1 := int(num1i64)
		num2i64, err := strconv.ParseInt(nums[i+1], 10, 64)
		if err != nil {
			panic(err)
		}
		num2 := int(num2i64)
		if _, ok := rules[num1]; !ok {
			panic(fmt.Sprintf("No num1 %d not found", num1))
		}
		if _, ok := rules[num2]; !ok {
			panic(fmt.Sprintf("No num2 %d not found", num2))
		}
		if rules[num2] <= rules[num1] {
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

func solvePart1(pageRules, updatedPages string) int {
	totMids := 0
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
		out := make([]int, len(rules))
		for k, v := range rules {
			out[v] = k
		}
		if correctOrder {
			midNum := printNums[len(printNums)/2]
			totMids += int(midNum)
		}
		// for i, rule := range lines {

		// 	left, right, err := parseRule(rule)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	pLeft := rules[left]
		// 	pRight := rules[right]

		// if pLeft == pRight {
		// 	for num, pos := range rules {
		// 		if pos > pRight {
		// 			rules[num]++
		// 		}
		// 	}
		// 	rules[right] = pRight + 1
		// } else if pLeft > pRight {
		// 	countZeros := 0
		// 	for _, v := range rules {
		// 		if v == 0 {
		// 			countZeros++
		// 		}
		// 	}

		// 	if pRight == 0 && countZeros > 1 {
		// 		for num, pos := range rules {
		// 			if pos > pLeft {
		// 				rules[num]++
		// 			}
		// 		}
		// 		rules[right] = pLeft + 1
		// 	} else {
		// 		for num, pos := range rules {
		// 			if pRight <= pos && pos <= pLeft {
		// 				rules[num]++
		// 			}
		// 		}
		// 		rules[left] = pRight
		// 	}
		// }

		// if !foundLeft && !foundRight {
		// 	rules[left] = 0
		// 	rules[right] = len(rules) - 1
		// } else if foundLeft && !foundRight {
		// 	if rules[left] < len(rules)-1 {
		// 		rules[right] = rules[left] + 1
		// 	} else {
		// 		rules[left] -= 1
		// 		rules[right] = rules[left] + 1
		// 		// fmt.Println("Tight snug to right")
		// 	}
		// } else if !foundLeft && foundRight {
		// 	if rules[right] > 0 {
		// 		rules[left] = rules[right] - 1
		// 	} else {
		// 		for num, pos := range rules {
		// 			if pos >= rules[right] && pos < len(rules)-1 {
		// 				rules[num] += 1
		// 			}
		// 		}
		// 		fmt.Println("new scenario", left, right)

		// 	}
		// 	// if
		// 	// rules[right] = rules[left] + 1

		// } else {
		// 	if rules[right] <= rules[left] {
		// 		diff := int(math.Abs(float64(rules[right] - rules[left])))
		// 		for num, pos := range rules {
		// 			if num != left && num != right && pos >= rules[right] {
		// 				rules[num] = pos + diff + 1
		// 			}
		// 		}
		// 		rules[right] = rules[right] + 1 + diff
		// 	}
		// }
		// if i+1 >= 1006 && i+1 < 1010 {
		// 	out := make([]int, len(rules))
		// 	for k, v := range rules {
		// 		if v != -1 {
		// 			out[v] = k
		// 		}
		// 	}
		// 	fmt.Println(i+1, out)
		// }
		// }
	}
	// numLines := strings.Split(updatedPages, "\n")
	// for _, line := range numLines {
	// 	ordered, midNum := checkLines(rules, line)
	// 	if ordered {
	// 		fmt.Println("good line", line)
	// 		sumMids += midNum
	// 	}
	// }
	return totMids
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

	// 	pageRules = `47|53
	// 97|13
	// 97|61
	// 97|47
	// 53|29
	// 75|29
	// 61|13
	// 75|53
	// 29|13
	// 97|29
	// 53|29
	// 61|53
	// 97|53
	// 61|29
	// 47|13
	// 75|47
	// 97|75
	// 47|61
	// 75|61
	// 47|29
	// 75|13
	// 53|13`

	// 	updatedPages = `75,47,61,53,29
	// 97,61,53,29,13
	// 75,29,13
	// 75,97,47,61,53
	// 61,13,29
	// 97,13,75,29,47`

	// 	updatedPages = `79,99,31,72,34,11,43,42,39,52,32,77,93
	// 34,88,46,72,93,56,63
	// 15,12,99,92,84,48,76,82,79,81,61,31,56,63,46,45,72
	// 46,79,54,76,27,37,96
	// 42,34,79,72,54,76,48`

	res := solvePart1(pageRules, updatedPages)
	utils.PrintSolution(5, 1, res)
}
