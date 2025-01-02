package day23

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestFindCliques(t *testing.T) {
	lines := []string{
		"S-1",
		"S-2",
		"S-3",
		"S-4",
		"1-2",
		"1-5",
		"2-3",
		"2-5",
		"3-4",
		"3-5",
	}
	g := buildGraph(lines)
	arg := "S"
	// got := g.findCliques(arg, "", 3, arg, map[[2]string]bool{})
	got := g.findCliques(arg, 3, map[string]bool{})
	want := 3
	if got != want {
		t.Errorf("findCliques(startNode=%s, cliqueSize=%d)=%d want %d", arg, 3, got, want)
	}
}

func TestSolvePart(t *testing.T) {

	lines, err := utils.ReadLines("input_test.txt")
	if err != nil {
		panic(err)
	}
	got := solvePart1(lines)
	want := 7
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
