package day20

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestSolveDay1(t *testing.T) {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		panic(err)
	}
	got := solvePart1(lines, 100)
	want := 1490
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolveDay2_mini(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")

	if err != nil {
		panic(err)
	}
	got := solvePart2(lines, 50, 20)
	want := 32 + 31 + 29 + 39 + 25 + 23 + 20 + 19 + 12 + 14 + 12 + 22 + 4 + 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolveDay2_part2(t *testing.T) {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		panic(err)
	}
	got := solvePart2(lines, 100, 2)
	want := 1490
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
