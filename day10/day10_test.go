package day10

import (
	"strings"
	"testing"
)

func pTestPart1(t *testing.T) {
	input := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
	lines := strings.Split(input, "\n")
	got := solvePart1(lines)
	want := 2
	if want != got {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}

	input = `...0...
...1...
...2...
6443456
7.....7
8.....8
9.....9`
	lines = strings.Split(input, "\n")
	got = solvePart1(lines)
	want = 1
	if want != got {
		t.Errorf("solverPart1 got %d want %d", got, want)
	}

	input = `...0...
...1...
...3...
6543456
7.....7
8.....8
9.....9`
	lines = strings.Split(input, "\n")
	got = solvePart1(lines)
	want = 0
	if want != got {
		t.Errorf("solverPart1 got %d want %d", got, want)
	}

	input = `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	lines = strings.Split(input, "\n")
	got = solvePart1(lines)
	want = 4
	if want != got {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}

	input = `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`
	lines = strings.Split(input, "\n")
	got = solvePart1(lines)
	want = 3
	if want != got {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}

	input = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	lines = strings.Split(input, "\n")
	got = solvePart1(lines)
	want = 36
	if want != got {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`
	lines := strings.Split(input, "\n")
	got := solvePart2(lines)
	want := 3
	if want != got {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}
	//	input = `..90..9
	//
	// ...1.98
	// ...2..7
	// 6543456
	// 765.987
	// 876....
	// 987....`
	//
	//	lines = strings.Split(input, "\n")
	//	got = solvePart1(lines)
	//	want = 4
	//	if want != got {
	//		t.Errorf("solvePart1 got %d want %d", got, want)
	//	}
}
