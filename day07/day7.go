package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type operator uint8

const (
	ADD operator = iota
	MULTIPLY
	CONCAT
	MAX_OP
)

func perm(combos *[][]operator, v []operator, options operator, n int, d int) {
	if d >= n {
		return
	}
	for opt := range options {
		v[d] = opt
		if n-1 == d {
			tmp := make([]operator, len(v))
			copy(tmp, v)
			*combos = append(*combos, tmp)
		}
		perm(combos, v, options, n, d+1)
	}

}

func operatorCombos(n int, max_operator operator) [][]operator {
	combos := [][]operator{}
	holder := make([]operator, n)
	perm(&combos, holder, max_operator, n, 0)
	return combos
}

func parseLine(line string) (int, []int, error) {
	lineSplit := strings.Split(line, ":")
	resultI64, err := strconv.ParseInt(lineSplit[0], 10, 64)
	if err != nil {
		return 0, nil, err
	}
	numsStr := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
	nums := make([]int, len(numsStr))
	for i := range len(numsStr) {
		numI64, err := strconv.ParseInt(numsStr[i], 10, 64)
		if err != nil {
			return 0, nil, err
		}
		nums[i] = int(numI64)
	}
	result := int(resultI64)
	return result, nums, nil
}

func solvePart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		result, numbers, err := parseLine(line)
		if err != nil {
			panic(err)
		}

		permuations := operatorCombos(len(numbers)-1, CONCAT)
		checkOut := false
		for _, perm := range permuations {
			// check permutation
			interimRes := numbers[0]
			for i, op := range perm {
				switch op {
				case ADD:
					interimRes = interimRes + numbers[i+1]
				case MULTIPLY:
					interimRes = interimRes * numbers[i+1]
				}
			}
			checkOut = interimRes == result
			if checkOut {
				break
			}
		}
		if checkOut {
			sum += result
		}

	}
	return sum
}

func solvePart2(lines []string) int {
	sum := 0
	for _, line := range lines {
		result, numbers, err := parseLine(line)
		if err != nil {
			panic(err)
		}

		permuations := operatorCombos(len(numbers)-1, MAX_OP)

		checkOut := false
		for _, perm := range permuations {
			// check permutation
			interimRes := numbers[0]
			for i, op := range perm {
				switch op {
				case ADD:
					interimRes = interimRes + numbers[i+1]
				case MULTIPLY:
					interimRes = interimRes * numbers[i+1]
				case CONCAT:
					conactStr := fmt.Sprintf("%d%d", interimRes, numbers[i+1])
					interimResI64, err := strconv.ParseInt(conactStr, 10, 64)
					if err != nil {
						panic(err)
					}
					interimRes = int(interimResI64)
				}
			}
			checkOut = interimRes == result
			if checkOut {
				break
			}
		}
		if checkOut {
			sum += result
		}

	}
	return sum
}

func Solve() {
	lines, err := utils.ReadLines("day07/input")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(7, 1, res)
	res = solvePart2(lines)
	utils.PrintSolution(7, 2, res)

}
