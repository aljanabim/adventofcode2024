package day12

import (
	"strings"
	"testing"
)

func TestPart1_simple(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	lines := strings.Split(input, "\n")
	want := 140
	got := solvePart1(lines)
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}
}

func TestPart1_xo(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	lines := strings.Split(input, "\n")
	want := 772
	got := solvePart1(lines)
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}
}

func TestPart1_big(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	lines := strings.Split(input, "\n")
	want := 1930
	got := solvePart1(lines)
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}
}
