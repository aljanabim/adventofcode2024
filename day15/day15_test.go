package day15

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestComputeGPSSum(t *testing.T) {
	input := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`

	lines := strings.Split(input, "\n")
	warehouse := readMap(lines)
	want := 10092
	got := computeGPSSum(&warehouse)
	if got != want {
		t.Errorf("computeGPSSum got %d want %d", got, want)

	}
}

func TestComputeGPSSum_mini(t *testing.T) {
	input := `########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`

	lines := strings.Split(input, "\n")
	warehouse := readMap(lines)
	want := 2028
	got := computeGPSSum(&warehouse)
	if got != want {
		t.Errorf("computeGPSSum got %d want %d", got, want)

	}
}

func TestMove_right(t *testing.T) {
	input1 := `@.OO.O.#`
	w := readMap([]string{input1})

	got := w.Grid
	want := [][]Entity{{3, 0, 1, 1, 0, 1, 0, 2}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}
	moves := []Dir{RIGHT}
	w.Step(moves)

	got = w.Grid
	want = [][]Entity{{0, 3, 1, 1, 0, 1, 0, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 0, 3, 1, 1, 1, 0, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 0, 0, 3, 1, 1, 1, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 0, 0, 3, 1, 1, 1, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}

	moves = []Dir{LEFT, LEFT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 3, 0, 0, 1, 1, 1, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}
}

func TestMove_full(t *testing.T) {
	rawData, err := os.ReadFile("map_mini.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawData), "\n")
	warehouse := readMap(lines)

	moves, err := readMoves("moves_mini.txt")
	if err != nil {
		panic(err)
	}

	warehouse.Step(moves)
	got := computeGPSSum(&warehouse)
	want := 10092
	if got != want {
		t.Fatalf("Move full: got %d want %d", got, want)
	}

}
