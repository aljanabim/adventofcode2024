package day22

import (
	"strconv"

	"github.com/aljanabim/adventofcode2024/utils"
)

func generateSecret(initialSecret int, nSecrets int) int {
	secret := initialSecret
	for range nSecrets {
		secret ^= secret * 64
		secret %= 16777216
		secret ^= secret / 32
		secret %= 16777216
		secret ^= secret * 2048
		secret %= 16777216
	}
	return secret
}

func solvePart1(lines []string) int {
	res := 0
	for _, line := range lines {
		nI64, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		initialSecret := int(nI64)
		res += generateSecret(initialSecret, 2000)
	}
	return res
}
func Solve() {
	lines, err := utils.ReadLines("day22/input.txt")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(22, 1, res)
}
