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
	got := len(g.findCliques(arg, 3, map[string]bool{}))
	want := 3
	if got != want {
		t.Errorf("findCliques(startNode=%s, cliqueSize=%d)=%d want %d", arg, 3, got, want)
	}
}

func TestSolvePart1(t *testing.T) {
	lines, err := utils.ReadLines("input_test.txt")
	if err != nil {
		panic(err)
	}
	graph := buildGraph(lines)
	got := solvePart1(graph)
	want := 7
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	lines, err := utils.ReadLines("input_test.txt")
	if err != nil {
		panic(err)
	}
	graph := buildGraph(lines)
	got := solvePart2(graph)
	want := "co,de,ka,ta"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
