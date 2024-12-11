package day10

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
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
		t.Errorf("solvePart1 got %d %d want", got, want)
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
		t.Errorf("solverPart1 got %d %d want", got, want)
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
		t.Errorf("solverPart1 got %d %d want", got, want)
	}

	/*
	   ..90..9
	   ...1.98
	   ...2..7
	   6543456
	   765.987
	   876....
	   987....
	*/

	input = `..90...
...1...
...2...
6543456
76.....
8......
9......`
	lines = strings.Split(input, "\n")
	got = solvePart1(lines)
	want = 1
	if want != got {
		t.Errorf("solvePart1 got %d %d want", got, want)
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
	//		t.Errorf("solvePart1 got %d %d want", got, want)
	//	}
}
