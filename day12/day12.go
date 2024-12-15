package day12

import (
	"fmt"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Region struct {
	area      int
	perimeter int
}

func checkPerimeter(name byte, row, col int, lines *[]string) (int, int, int, int) {
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
	top, right, left, bottom := checkPerimeter(name, row, col, lines)
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

func visitNeighbors2(name byte, row, col, sourceRow, sourceCol int, lines *[]string, cache Cache, checkedRows, checkedCols map[int]bool, sourceHas sides) (int, int) {
	cache[[2]int{row, col}] = true
	top, right, bottom, left := checkPerimeter(name, row, col, lines)
	newSourceHas := sides{itob(top), itob(right), itob(bottom), itob(left)}

	countPerimeter := sides{itob(top), itob(right), itob(bottom), itob(left)}
	horzDir := col - sourceCol // 1 = left to right | -1 = right to left
	vertDir := row - sourceRow // 1 = top to bottom | -1 = bottom to top

	if sourceHas.top && horzDir != 0 {
		countPerimeter.top = false
	}
	if sourceHas.bottom && horzDir != 0 {
		countPerimeter.bottom = false
	}
	if sourceHas.right && vertDir != 0 {
		countPerimeter.right = false
	}
	if sourceHas.left && vertDir != 0 {
		countPerimeter.left = false
	}
	// going right to left with a neighbor means we should check if neighbor above has a left
	if horzDir == -1 && !itob(top) {
		coutLeft := false
		neiName := name
		// if row > 0 && (*lines)[row-1][col] == name {
		// 	_, _, _, pl := checkPerimeter(name, row-1, col, lines)
		// 	fmt.Println("checking guy above", pl, cache[[2]int{row - 1, col}])
		// 	if itob(pl) && !cache[[2]int{row - 1, col}] {
		// 		// coutLeft = true
		// 	}
		// }
		nrow := row - 1
		neiName = (*lines)[nrow][col]
		neiVisited := false
		anyNeiHasLeftPer := false

		for neiName == name && nrow >= 0 {
			neiName = (*lines)[nrow][col]
			if neiName == name {
				neiVisited = neiVisited || cache[[2]int{nrow, col}]
				fmt.Println("looping in left", string(neiName), nrow)
				_, _, _, lp := checkPerimeter(name, nrow, col, lines)
				anyNeiHasLeftPer = anyNeiHasLeftPer || itob(lp)
			}
			// 	if neiName != name {
			// 		coutLeft = true
			// 		break
			// 	}
			nrow--
		}
		fmt.Println("Any have been visited", neiVisited, anyNeiHasLeftPer)
		// _, _, _, nl := checkPerimeter(name, nrow, col, lines)
		// if itob(nl) { // above neighbor has no left side
		countPerimeter.left = coutLeft
		// }
	}

	// going down and row has already been checked (don't count bottom result if any)
	if vertDir == 1 && checkedRows[row] {
		fmt.Println("checking naught line")
		if col < len((*lines)[0])-1 { // check right neighbor
			nei := (*lines)[row][col+1]
			if nei == name {
				_, _, nb, _ := checkPerimeter(name, row, col+1, lines)
				if itob(nb) { // if neighbor has bottom skip our own bottom
					fmt.Println("checking neightbor", nb)
					countPerimeter.bottom = false
				}
			}
		}
		if col > 0 { // check left neighbor
			nei := (*lines)[row][col-1]
			if nei == name {
				_, _, nb, _ := checkPerimeter(name, row, col-1, lines)
				if itob(nb) { // if neighbor has bottom skip our own bottom
					fmt.Println("checking neightbor", nb)
					countPerimeter.bottom = false
				}
			}
		}
	}
	checkedRows[row] = row == sourceRow || vertDir == 1

	// going up and previous cell had no bottom, then we must reset the previous row for until entered again when going down
	if vertDir == -1 && !sourceHas.bottom {
		checkedRows[sourceRow] = false
	}

	// checkedCols[col] = col == sourceCol

	// if row > sourceRow { // came from top
	// 	newSourcePerimeter.top = false
	// } else if row < sourceRow { // came from bottom
	// 	newSourcePerimeter.bottom = false
	// }

	// above
	// newSourcePerimeter := sides{top: !checkedRows[row-1], right: !itob(right), bottom: false, left: !itob(left)}
	// right
	// newCountPerimeter := sides{top: !itob(top), right: !checkedCols[col+1], bottom: !itob(bottom), left: false}
	// below
	//newCountPerimeter := sides{top: false, right: !itob(right), bottom: !checkedRows[row+1], left: !itob(left)}
	// left
	// newCountPerimeter := sides{top: !itob(top), right: false, bottom: !itob(bottom), left: !checkedCols[col-1]}

	// Below is logic to cover situations where we should override the checkedRows and checkedCols conditions

	// check if solo corner on the left (the only way to enter such an edge is if visited from right so count top and count bottom
	// should be covered by the logic in the recursive calls)
	// if itob(top) && itob(left) && itob(bottom) {
	// countPerimeter.top = true
	// countPerimeter.left = true
	// countPerimeter.bottom = true
	// }
	// check if solo corner on the right (the only way to enter such an edge is if visited from left so count top and count bottom
	// should be covered by the logic in the recursive calls)
	// if itob(top) && itob(right) && itob(bottom) {
	// countPerimeter.top = true
	// countPerimeter.right = true
	// countPerimeter.bottom = true
	// }
	// check if solo corner on the top (the only way to enter such an edge is if visited from bottom so count right and count right
	// should be covered by the logic in the recursive calls)
	// if itob(left) && itob(top) && itob(right) {
	// countPerimeter.top = true
	// countPerimeter.right = true
	// countPerimeter.left = true
	// }
	// check if solo corner on the bottom (the only way to enter such an edge is if visited from top so count right and count right
	// should be covered by the logic in the recursive calls)
	// if itob(left) && itob(bottom) && itob(right) {
	// countPerimeter.right = true
	// countPerimeter.left = true
	// countPerimeter.bottom = true
	// }

	// if itob(top) && itob(left) { // check top left corner (no need to check top right)
	// 	countPerimeter.top = true
	// 	countPerimeter.left = true
	// }
	// if itob(bottom) && itob(right) { // check bottom right (no need to check top left given order of visiting neighbors [top, right, bottom, left])
	// countPerimeter.right = true
	// countPerimeter.bottom = true
	// }

	/*
		RRRRIICCFF
		RRRRIICCCF
		VVRRRCCFFF
		VVRCCCJFFF
		VVVVCJJCFE
		VVIVCCJJEE
		VVIIICJJEE
		MIIIIIJJEE
		MIIISIJEEE
		MMMISSJEEE

		XRRRRRX
		XRRRXRX
		XXRXXXX
		XXRXXXX
		XXXXXXX
	*/

	perimeter := top*btoi(countPerimeter.top) + right*btoi(countPerimeter.right) + bottom*btoi(countPerimeter.bottom) + left*btoi(countPerimeter.left)
	if name == 'X' || name == 'C' {
		fmt.Print("visiting \"", string(name), "\" at ", row, " ", col, " perim ", perimeter, "\t==\t")
		fmt.Print(top, countPerimeter.top, right, countPerimeter.right, bottom, countPerimeter.bottom, left, countPerimeter.left)
		fmt.Println("\tsource has", sourceHas)
	}
	area := 1

	// check above
	if row > 0 && (*lines)[row-1][col] == name {
		if visited := cache[[2]int{row - 1, col}]; !visited {
			// newCountPerimeter := sides{top: !checkedRows[row-1], right: !itob(right), bottom: false, left: !itob(left)}
			a, p := visitNeighbors2(name, row-1, col, row, col, lines, cache, checkedRows, checkedCols, newSourceHas)
			area += a
			perimeter += p
		}
	}

	// check right
	if col < len((*lines)[0])-1 && (*lines)[row][col+1] == name {
		if visited := cache[[2]int{row, col + 1}]; !visited {
			// newCountPerimeter := sides{top: !itob(top), right: !checkedCols[col+1], bottom: !itob(bottom), left: false}
			a, p := visitNeighbors2(name, row, col+1, row, col, lines, cache, checkedRows, checkedCols, newSourceHas)
			area += a
			perimeter += p
		}
	}

	// check below
	if row < len(*lines)-1 && (*lines)[row+1][col] == name {
		if visited := cache[[2]int{row + 1, col}]; !visited {
			// newCountPerimeter := sides{top: false, right: !itob(right), bottom: !checkedRows[row+1], left: !itob(left)}
			a, p := visitNeighbors2(name, row+1, col, row, col, lines, cache, checkedRows, checkedCols, newSourceHas)
			area += a
			perimeter += p
		}
	}

	// check left
	if col > 0 && (*lines)[row][col-1] == name {
		if visited := cache[[2]int{row, col - 1}]; !visited {
			// newCountPerimeter := sides{top: !itob(top), right: false, bottom: !itob(bottom), left: !checkedCols[col-1]}
			a, p := visitNeighbors2(name, row, col-1, row, col, lines, cache, checkedRows, checkedCols, newSourceHas)
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
		fmt.Println(region.area, "x", region.perimeter, "=", region.area*region.perimeter)
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

func solvePart2(lines []string) int {
	visitCache := make(Cache)
	regions := []Region{}

	for row, line := range lines {
		for col, name := range line {
			if visited := visitCache[[2]int{row, col}]; !visited {
				checkedRows := map[int]bool{}
				checkedCols := map[int]bool{}
				area, perimeter := visitNeighbors2(byte(name), row, col, row, col, &lines, visitCache, checkedRows, checkedCols, sides{})
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
	// 1.359.028 - upper limit
	//   821.799 - too low
	//   835.777 - too low
}
