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
	got := solvePart1(lines)
	want := 1490
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
