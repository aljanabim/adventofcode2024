package day06

import (
	"fmt"
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

func TestSolvePart2_first(t *testing.T) {
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

	got, newObs := solvePart2(gridObstacles, currPos, rows, cols)
	for obs, _ := range newObs {
		if obs[0] >= 0 && obs[0] < len(lines) {
			lines[obs[0]] = lines[obs[0]][:obs[1]] + "O" + lines[obs[0]][obs[1]+1:]
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	want := 6
	if got != want {
		t.Errorf("solvePart2 want %d got %d", want, got)
	}
}

func TestSolvePart2_second(t *testing.T) {
	linesTxt := `....#.....
.........#
..........
..........
.#........
........#.
....^.....
..........
..........
..........`

	linesTxt = `....#.....
.........#
..........
...#......
.......#..
..........
.#..^.....
........#.
..........
..........`

	/*
	   	linesTxt = `##..
	   ...#
	   ....
	   ^.#.` // expects 0

	   	linesTxt = `.#...
	   ....#
	   .....
	   .^.#.
	   #....
	   ..#..
	   ` // expects 3
	*/

	lines := strings.Split(linesTxt, "\n")
	rows, cols, gridObstacles, currPos := buildGrid(lines)

	got, newObs := solvePart2(gridObstacles, currPos, rows, cols)
	for obs, _ := range newObs {
		if obs[0] >= 0 && obs[0] < len(lines) {
			lines[obs[0]] = lines[obs[0]][:obs[1]] + "O" + lines[obs[0]][obs[1]+1:]
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	want := 4
	if got != want {
		t.Errorf("solvePart2 want %d got %d", want, got)
	}
}
