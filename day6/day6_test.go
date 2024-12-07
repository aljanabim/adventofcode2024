package day6

import (
	"strings"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	linesTxt := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	lines := strings.Split(linesTxt, "\n")
	rows, cols, gridObstacles, currPos := buildGrid(lines)

	got := solvePart1(gridObstacles, currPos, rows, cols)
	want := 41
	if got != want {
		t.Fatalf("solvePart1 want %d got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	linesTxt := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	lines := strings.Split(linesTxt, "\n")
	rows, cols, gridObstacles, currPos := buildGrid(lines)

	got := solvePart2(gridObstacles, currPos, rows, cols)
	want := 6
	if got != want {
		t.Fatalf("solvePart2 want %d got %d", want, got)
	}
}

/*
....#.....
....xxxxx#
....x...x.
..#.x...x.
..xxxxx#x.
..x.x.x.x.
.#xxxxxxx.
.xxxxxxx#.
#xxxxxxx..
......#x..

....#.....
....XXXXX#
....X...X.
..#.X...X.
..XXXXX#X.
..X.X.X.X.
.#XoXXXXX.
.XXXXXoo#.
#oXoXXXX..
......#o..
*/
