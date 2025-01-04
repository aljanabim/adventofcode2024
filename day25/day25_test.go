package day25

import "testing"

func TestSolvePart1(t *testing.T) {
	locks, keys, err := parseInput("input_test.txt")
	if err != nil {
		panic(err)
	}
	got := solvePart1(locks, keys)
	want := 3
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
