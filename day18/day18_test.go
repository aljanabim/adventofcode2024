package day18

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestSearchPath_empty(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}
	got := solvePart1(lines[:0], 7)
	want := 12
	if got != want {
		t.Fatalf("SearchPath got %d want %d", got, want)
	}

}

func TestSearchPath(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}
	got := solvePart1(lines[:12], 7)
	want := 22
	if got != want {
		t.Fatalf("SearchPath got %d want %d", got, want)
	}

}

func TestSolvePart2(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}
	got := solvePart2(lines, 7, 12)
	want := "6,1"
	if got != want {
		t.Fatalf("SearchPath got (%s) want (%s)", got, want)
	}

}
