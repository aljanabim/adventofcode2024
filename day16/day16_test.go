package day16

import (
	"strings"
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestSearch_tiny(t *testing.T) {
	input := `######
#...E#
##.#.#
##...#
##.#.#
##S..#
######`
	lines := strings.Split(input, "\n")
	maze := createMaze(lines)
	got := Search(maze.Start, maze.Start, RIGHT, &maze, map[[2]int]bool{}, 0)
	want := 1005.0
	if got != want {
		t.Fatalf("got %f want %f", got, want)
	}
	/*
	  ######
	  #...E#
	  ##.#.#
	  ##...#
	  ##.#.#
	  ##S..#
	  ######
	*/
}

func TestSearch_mini(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}

	maze := createMaze(lines)
	got := Search(maze.Start, maze.Start, RIGHT, &maze, map[[2]int]bool{}, 0)

	want := 2.0
	if got != want {
		t.Fatalf("got %f want %f", got, want)
	}
}

func TestSearch_mini3(t *testing.T) {
	lines, err := utils.ReadLines("input_mini3.txt")
	if err != nil {
		panic(err)
	}

	maze := createMaze(lines)
	got := Search(maze.Start, maze.Start, RIGHT, &maze, map[[2]int]bool{}, 0)

	want := 2.0
	if got != want {
		t.Fatalf("got %f want %f", got, want)
	}
}
