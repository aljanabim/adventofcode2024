package day19

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestSolvePart1_simple(t *testing.T) {
	stripes := []int{0, 22, 3, 12}
	want := 312
	maxOrder := getOrder(want)
	got := helper(stripes, want, 0, maxOrder, 0)
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}

}

func _TestSolvePart1_mini(t *testing.T) {
	stripes := []int{0, 1, 21, 3, 4, 325, 13, 43, 31}
	cases := []struct {
		arg  int
		want int
	}{
		{31211, 31211},
		{3441, 3441},
		{4331, 4331},
		{113431, 113431},
		{5325, -1},
		{325114, 325114},
		{3141, 3141},
		{331423, -1},
	}
	for _, c := range cases {
		got := helper(stripes, c.arg, 0, getOrder(c.arg), 0)
		if got != c.want {
			t.Errorf("solvePart1(%d)=%d want %d", c.arg, got, c.want)
		}
	}

}

func _TestSolvePart1(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}
	stripesStr := lines[0]
	flags := lines[2:]
	stripes, color2Int := parseStripes(stripesStr)
	flagNums := parseFlags(flags, color2Int)

	got := solvePart1(stripes, flagNums)
	want := 6
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
