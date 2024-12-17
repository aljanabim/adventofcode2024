package day14

import (
	"reflect"
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestParseLines(t *testing.T) {
	lines, err := utils.ReadLines("input_test_parse.txt")
	if err != nil {
		panic(err)
	}
	want := []Robot{{X: 0, Y: 4, Vx: 3, Vy: -3}, {X: 6, Y: 3, Vx: -1, Vy: -3}, {X: 10, Y: 3, Vx: -1, Vy: 2}}
	got := parseLines(lines)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("parseLines \nwant %+v \ngot %+v", want, got)
	}
}

func TestSimulate(t *testing.T) {
	n := 5
	robot := Robot{X: 2, Y: 4, Vx: 2, Vy: -3}
	robot.Step(n, 11, 7)
	want := Robot{X: 1, Y: 3, Vx: 2, Vy: -3}
	if !reflect.DeepEqual(want, robot) {
		t.Fatalf("robot stepping %d steps \ngot %+v \nwant %+v", n, robot, want)
	}
}

func TestSolvePart1(t *testing.T) {
	lines, err := utils.ReadLines("input_test.txt")
	if err != nil {
		panic(err)
	}
	got := solvePart1(lines, 100, 11, 7)
	want := 12
	if got != want {
		t.Fatalf("solveDay1 got %d want %d", got, want)
	}
}
