package day12

import "github.com/aljanabim/adventofcode2024/utils"

type Region struct {
	area      int
	perimeter int
}

func computePerimeter(name byte, row, col int, lines *[]string) int {
	var neighbor byte
	perimeter := 0

	// check top row and bottom row
	if row == 0 || row == len(*lines)-1 {
		perimeter++
	}

	// check above
	if row > 0 {
		neighbor = (*lines)[row-1][col]
		if neighbor != name {
			perimeter++
		}
	}

	// check below
	if row < len(*lines)-1 {
		neighbor = (*lines)[row+1][col]
		if neighbor != name {
			perimeter++
		}
	}

	// check left or right
	if col == 0 || col == len((*lines)[0])-1 {
		perimeter++
	}

	// check left
	if col > 0 {
		neighbor = (*lines)[row][col-1]
		if neighbor != name {
			perimeter++
		}
	}

	// check right
	if col < len((*lines)[0])-1 {
		neighbor = (*lines)[row][col+1]
		if neighbor != name {
			perimeter++
		}
	}

	return perimeter
}

/*
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
*/
func visitNeighbors(name byte, row, col int, lines *[]string, cache Cache) (int, int) {
	// fmt.Println("visiting", string(name), "at", row, col)
	cache[[2]int{row, col}] = true
	perimeter := computePerimeter(name, row, col, lines)
	area := 1

	// check above
	if row > 0 && (*lines)[row-1][col] == name {
		if visited := cache[[2]int{row - 1, col}]; !visited {
			a, p := visitNeighbors(name, row-1, col, lines, cache)
			area += a
			perimeter += p
		}
	}

	// check below
	if row < len(*lines)-1 && (*lines)[row+1][col] == name {
		if visited := cache[[2]int{row + 1, col}]; !visited {
			a, p := visitNeighbors(name, row+1, col, lines, cache)
			area += a
			perimeter += p
		}
	}

	// check left
	if col > 0 && (*lines)[row][col-1] == name {
		if visited := cache[[2]int{row, col - 1}]; !visited {
			a, p := visitNeighbors(name, row, col-1, lines, cache)
			area += a
			perimeter += p
		}
	}

	// check right
	if col < len((*lines)[0])-1 && (*lines)[row][col+1] == name {
		if visited := cache[[2]int{row, col + 1}]; !visited {
			a, p := visitNeighbors(name, row, col+1, lines, cache)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

type Cache map[[2]int]bool

func computeFenceCost(regions []Region) int {
	cost := 0
	for _, region := range regions {
		cost += region.area * region.perimeter
	}
	return cost
}

func solvePart1(lines []string) int {
	visitCache := make(Cache)
	regions := []Region{}

	regionIdx := -1
	for row, line := range lines {
		for col, name := range line {
			if visited := visitCache[[2]int{row, col}]; !visited {
				regionIdx++
				area, perimeter := visitNeighbors(byte(name), row, col, &lines, visitCache)
				regions = append(regions, Region{area, perimeter})
			}
		}
	}

	return computeFenceCost(regions)
}

func Solve() {
	lines, err := utils.ReadLines("day12/input")
	if err != nil {
		panic(err)
	}
	res := solvePart1(lines)
	utils.PrintSolution(12, 1, res)

}
