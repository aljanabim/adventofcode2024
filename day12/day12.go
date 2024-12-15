package day12

import (
	"github.com/aljanabim/adventofcode2024/utils"
)

type Region struct {
	area      int
	perimeter int
}

func computePerimeter(name byte, row, col int, lines *[]string) (int, int, int, int) {
	var neighbor byte
	topPerimeter := 0
	rightPerimeter := 0
	bottomPerimeter := 0
	leftPerimeter := 0

	// check top row and bottom row
	if row == 0 {
		topPerimeter++
	}
	if row == len(*lines)-1 {
		bottomPerimeter++
	}

	// check above
	if row > 0 {
		neighbor = (*lines)[row-1][col]
		if neighbor != name {
			topPerimeter++
		}
	}

	// check below
	if row < len(*lines)-1 {
		neighbor = (*lines)[row+1][col]
		if neighbor != name {
			bottomPerimeter++
		}
	}

	// check far left or far right
	if col == 0 {
		leftPerimeter++
	}
	if col == len((*lines)[0])-1 {
		rightPerimeter++
	}

	// check left
	if col > 0 {
		neighbor = (*lines)[row][col-1]
		if neighbor != name {
			leftPerimeter++
		}
	}

	// check right
	if col < len((*lines)[0])-1 {
		neighbor = (*lines)[row][col+1]
		if neighbor != name {
			rightPerimeter++
		}
	}

	return topPerimeter, rightPerimeter, bottomPerimeter, leftPerimeter
}

func visitNeighbors(name byte, row, col int, lines *[]string, cache Cache) (int, int) {
	cache[[2]int{row, col}] = true
	top, right, left, bottom := computePerimeter(name, row, col, lines)
	perimeter := top + right + left + bottom
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

/*
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
*/

func btoi(c bool) int {
	if c {
		return 1
	}
	return 0
}
func itob(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

type sides struct {
	top    bool
	right  bool
	bottom bool
	left   bool
}

/*
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
*/

func buildRegionGrid(grid *[][]bool, name byte, row, col int, lines *[]string, cache Cache) int {
	cache[[2]int{row, col}] = true
	(*grid)[row][col] = true
	area := 1

	// check above
	if row > 0 && (*lines)[row-1][col] == name {
		if visited := cache[[2]int{row - 1, col}]; !visited {
			a := buildRegionGrid(grid, name, row-1, col, lines, cache)
			area += a
		}
	}

	// check below
	if row < len(*lines)-1 && (*lines)[row+1][col] == name {
		if visited := cache[[2]int{row + 1, col}]; !visited {
			a := buildRegionGrid(grid, name, row+1, col, lines, cache)
			area += a
		}
	}

	// check left
	if col > 0 && (*lines)[row][col-1] == name {
		if visited := cache[[2]int{row, col - 1}]; !visited {
			a := buildRegionGrid(grid, name, row, col-1, lines, cache)
			area += a
		}
	}

	// check right
	if col < len((*lines)[0])-1 && (*lines)[row][col+1] == name {
		if visited := cache[[2]int{row, col + 1}]; !visited {
			a := buildRegionGrid(grid, name, row, col+1, lines, cache)
			area += a
		}
	}

	return area
}

type Cache map[[2]int]bool

func computeFenceCost(regions []Region) int {
	cost := 0
	for _, region := range regions {
		cost += region.area * region.perimeter
		// fmt.Println(region.area, "x", region.perimeter, "=", region.area*region.perimeter)
	}
	return cost
}

func solvePart1(lines []string) int {
	visitCache := make(Cache)
	regions := []Region{}

	for row, line := range lines {
		for col, name := range line {
			if visited := visitCache[[2]int{row, col}]; !visited {
				area, perimeter := visitNeighbors(byte(name), row, col, &lines, visitCache)
				regions = append(regions, Region{area, perimeter})
			}
		}
	}

	return computeFenceCost(regions)
}

func resetGrid(rows, cols int) [][]bool {
	grid := make([][]bool, rows)
	for row := range rows {
		grid[row] = make([]bool, cols)
	}
	return grid
}

func countSides(grid *[][]bool) int {
	count := 0
	// for row := range len(*grid) {
	// for col := range len((*grid)[row]) {
	for row, line := range *grid {
		horzSides := 0
		vertSides := 0
		for col, in := range line {
			// top & bottom sides
			if row == 0 || row == len((*grid))-1 {
				if col == 0 && in {
					horzSides++
				}
				if col > 0 && in && !line[col-1] {
					horzSides++
				}
			}
			// left & right sides
			if col == 0 || col == len(line)-1 {
				if row == 0 && in {
					vertSides++
				}
				if row > 0 && in && !(*grid)[row-1][col] {
					vertSides++
				}
			}

			// horizontal sides on left line
			if col == 0 && row > 0 {
				if in != (*grid)[row-1][col] {
					horzSides++
				}
			}
			// vertical sides on top line
			if row == 0 && col > 0 {
				if in != (*grid)[row][col-1] {
					vertSides++
				}
			}

			// inner parts
			if 0 < row && 0 < col {
				upLeftIn := (*grid)[row-1][col-1]
				upIn := (*grid)[row-1][col]
				leftIn := (*grid)[row][col-1]

				// INNER POINT LINES
				/* horizontal line due to inner point
				OX
				OO<-
				*/
				if in && upLeftIn && leftIn && !upIn {
					horzSides++
				}
				/* horizontal & vertical lines to inner point
				.X
				XO<-
				*/
				if in && !upIn && !leftIn {
					horzSides++
					vertSides++
				}
				/* vertical line due to inner point
				OO
				XO<-
				*/
				if in && upLeftIn && !leftIn && upIn {
					vertSides++
				}

				// OUTER POINT LINES
				/* horizontal line due to outer point
				XO
				.X<-
				*/
				if !in && upIn && !upLeftIn {
					horzSides++
				}
				/* horizontal lines due to outer point
				OO
				OX<-
				*/
				if !in && upIn && upLeftIn && leftIn {
					horzSides++
				}
				/* vertical lines due to outer point
				XX
				OX<-
				*/
				if !in && !upIn && !upLeftIn && leftIn {
					vertSides++
				}
				/* vertical lines due to outer point
				.O
				OX<-
				*/
				if !in && upIn && leftIn {
					vertSides++
				}
			}
		}
		// fmt.Println("Sides", line, "horz", horzSides, "vert", vertSides)
		count += horzSides + vertSides
	}
	// }
	// fmt.Println()
	// }
	return count
}

func solvePart2(lines []string) int {
	visitCache := make(Cache)
	regions := []Region{}

	for row, line := range lines {
		for col, name := range line {
			if visited := visitCache[[2]int{row, col}]; !visited {
				grid := resetGrid(len(lines), len(line))
				area := buildRegionGrid(&grid, byte(name), row, col, &lines, visitCache)
				perimeter := countSides(&grid)
				// fmt.Println("Area", area, "perimeter", perimeter)
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
	res = solvePart2(lines)
	utils.PrintSolution(12, 2, res)
}
