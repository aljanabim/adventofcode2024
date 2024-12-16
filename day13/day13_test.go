package day13

import (
	"reflect"
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestInverse(t *testing.T) {
	input := Matrix{
		{97, 22},
		{34, 67},
	}
	got, _ := inverse(input)
	want := Matrix{
		{67.0 / 5751.0, -22.0 / 5751.0},
		{-34.0 / 5751.0, 97.0 / 5751.0},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Inverse: \ngot  %v \nwant %v", got, want)
	}
}
func TestInverse_zero_det(t *testing.T) {
	input := Matrix{
		{0, 1},
		{0, 1},
	}

	got, err := inverse(input)
	if err == nil {
		t.Fatalf("Expected and error got %v", got)
	}
}

func TestPart1(t *testing.T) {
	lines, err := utils.ReadLines("input_base")
	if err != nil {
		panic(err)
	}

	got := solveDay(lines, false)
	want := 480
	if got != want {
		t.Errorf("solveDay1 got %d want %d", got, want)
	}
}

func TestPart1_full(t *testing.T) {
	lines, err := utils.ReadLines("input")
	if err != nil {
		panic(err)
	}

	got := solveDay(lines, false)
	want := 32041
	if got != want {
		t.Errorf("solveDay1 got %d want %d", got, want)
	}
}

func TestPart2_full(t *testing.T) {
	lines, err := utils.ReadLines("input")
	if err != nil {
		panic(err)
	}

	got := solveDay(lines, true)
	want := 95843948914827
	if got != want {
		t.Errorf("solveDay2 got %d want %d", got, want)
	}

}
