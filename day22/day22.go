package day22

import (
	"strconv"

	"github.com/aljanabim/adventofcode2024/utils"
)

func generateSecret(secret int) int {
	secret ^= (secret << 6)
	secret %= 1 << 24
	secret ^= (secret >> 5)
	secret %= 1 << 24
	secret ^= (secret << 11)
	secret %= 1 << 24
	return secret
}

func solvePart1(lines []string) int {
	defer utils.Duration(utils.Track("Part 1"))
	res := 0
	for _, line := range lines {
		nI64, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		secret := int(nI64)
		for range 2000 {
			secret = generateSecret(secret)
		}
		res += secret
	}
	return res
}

func solvePart2(lines []string) int {
	defer utils.Duration(utils.Track("Part 2"))
	res := 0
	seqs := map[[4]int][]int{}
	for n, line := range lines {
		nI64, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		secret := int(nI64)
		diffs := [2000]int{}
		prices := [2000]int{}
		for i := range 2000 {
			newSecret := generateSecret(secret)
			diffs[i] = (newSecret % 10) - (secret % 10)
			prices[i] = (newSecret % 10)
			if i >= 4 {
				seq := [4]int{diffs[i-3], diffs[i-2], diffs[i-1], diffs[i]}
				if val := seqs[seq]; len(val) == n {
					seqs[seq] = append(seqs[[4]int{diffs[i-3], diffs[i-2], diffs[i-1], diffs[i]}], prices[i])
				}
			}
			secret = newSecret
		}
		// make sure the length of a sequence is equal to n
		for k, v := range seqs {
			if len(v) < n+1 {
				seqs[k] = append(seqs[k], 0)
			}
		}
		res += secret
	}

	maxPrice := 0
	for _, prices := range seqs {
		if len(prices) > 1 {
			tot := 0
			for _, price := range prices {
				tot += price
			}
			if tot > maxPrice {
				maxPrice = tot
			}
		}
	}
	return maxPrice
}
func Solve() {
	lines, err := utils.ReadLines("day22/input.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(22, 1, res)
	res = solvePart2(lines)
	utils.PrintSolution(22, 2, res)
}
