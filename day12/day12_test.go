package day12

import (
	"strings"
	"testing"
)

func _TestPart1_simple(t *testing.T) {
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

func _TestPart1_xo(t *testing.T) {
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

func _TestPart1_big(t *testing.T) {
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

func TestPart2_simple(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	lines := strings.Split(input, "\n")
	want := 80
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_xomini(t *testing.T) {

	input := `OOOOOO
OOXXOO
OOOOOO`
	lines := strings.Split(input, "\n")
	want := 136
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_snake(t *testing.T) {

	input := `X00
XX0
000
000
X0X`
	lines := strings.Split(input, "\n")
	want := 18 + 4 + 4 + 10*12
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}
func TestPart2_xo(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	lines := strings.Split(input, "\n")
	want := 436
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_E(t *testing.T) {
	input := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	lines := strings.Split(input, "\n")
	want := 236
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_corners(t *testing.T) {
	input := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

	// 	input := `OOXXOO
	// OOOOOO` // 88

	lines := strings.Split(input, "\n")
	want := 368
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_big(t *testing.T) {
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
	want := 1206
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_doubledeep(t *testing.T) {
	input := `XRRRRRX
XRRRXRX
XXRXXXX
XXRXXXX
XXXXXXX`

	lines := strings.Split(input, "\n")
	want := 11*12 + 24*16
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}

func TestPart2_lefty(t *testing.T) {
	input := `CC
XC
CC
CC`
	lines := strings.Split(input, "\n")
	want := 7*8 + 4*1
	got := solvePart2(lines)
	if got != want {
		t.Errorf("solvePart2 got %d want %d", got, want)
	}
}
