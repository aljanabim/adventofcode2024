package day9

import (
	"github.com/aljanabim/adventofcode2024/utils"
)

func solvePart1(line string) int {
	nums := []int{}
	freeMem := []int{}
	id := 0
	for i, n := range line {
		n := int(n - '0')
		// length of file
		if i%2 == 0 {
			for range n {
				nums = append(nums, id)
			}
			id++
		}
		// length of free
		if (i+1)%2 == 0 && i > 0 {
			for range n {
				freeMem = append(freeMem, len(nums))
				nums = append(nums, -1)
			}
		}
	}
	nextFreeId := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != -1 && nextFreeId < len(freeMem) && i > freeMem[nextFreeId] {
			nextFreeIdx := freeMem[nextFreeId]
			nums[nextFreeIdx] = nums[i]
			nums[i] = -1
			nextFreeId++
		}
	}
	// compute hash
	hash := 0
	for i, n := range nums {
		if n != -1 {
			hash += i * n
		}
	}
	return hash

}
func Solve() {
	input, err := utils.ReadLines("day9/input")
	if err != nil {
		panic(err)
	}
	res := solvePart1(input[0])
	utils.PrintSolution(9, 1, res)
}
