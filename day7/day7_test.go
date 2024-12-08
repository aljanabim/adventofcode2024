package day7

import (
	"strings"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	lines := strings.Split(input, "\n")
	want := 3749
	got := solvePart1(lines)
	if got != want {
		t.Errorf("solvePart1 want %d got %d", want, got)
	}

}

func TestSolvePart2(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	lines := strings.Split(input, "\n")
	want := 11387
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart1 want %d got %d", want, got)
	}

}
