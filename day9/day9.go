package day9

import (
	"github.com/aljanabim/adventofcode2024/utils"
)

func checksum(nums []int) int {
	// compute hash
	hash := 0
	for i, n := range nums {
		if n != -1 {
			hash += i * n
		}
	}
	return hash
}

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
	return checksum(nums)
}

type fileMem struct {
	id   int
	pos  int
	size int
}
type freeMem struct {
	pos  int
	size int
}

func solvePart2(line string) int {
	nums := []int{}
	files := []*fileMem{}
	freeMems := []*freeMem{}
	id := 0
	for i, n := range line {
		n := int(n - '0')
		// length of file
		if i%2 == 0 {
			files = append(files, &fileMem{id, len(nums), n})
			for range n {
				nums = append(nums, id)
			}
			id++
		}
		// length of free
		if (i+1)%2 == 0 && i > 0 {
			freeMems = append(freeMems, &freeMem{pos: len(nums), size: n})
			for range n {
				nums = append(nums, -1)
			}
		}
	}

	for f := len(files) - 1; f >= 0; f-- {
		file := files[f]
		for _, free := range freeMems {
			if file.pos > free.pos && free.size >= file.size {
				for i := range file.size {
					nums[free.pos+i] = file.id
					nums[file.pos+i] = -1
				}
				file.pos = free.pos
				free.size = free.size - file.size
				free.pos += file.size
				break
			}
		}
	}
	return checksum(nums)
}

func Solve() {
	input, err := utils.ReadLines("day9/input")
	if err != nil {
		panic(err)
	}
	res := solvePart1(input[0])
	utils.PrintSolution(9, 1, res)
	res = solvePart2(input[0])
	utils.PrintSolution(9, 2, res)
}
