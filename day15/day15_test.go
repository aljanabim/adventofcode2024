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

func TestComputeGPSSumExtended(t *testing.T) {
	input := `####################
##[].......[].[][]##
##[]...........[].##
##[]........[][][]##
##[]......[]....[]##
##..##......[]....##
##..[]............##
##..@......[].[][]##
##......[][]..[]..##
####################`

	lines := strings.Split(input, "\n")
	warehouse := readMap(lines)
	want := 9021
	got := computeGPSSum(&warehouse)
	if got != want {
		t.Errorf("computeGPSSumExtended got %d want %d", got, want)

	}
}

func TestMove_right(t *testing.T) {
	input1 := `@.OO.O.#`
	w := readMap([]string{input1})

	got := w.Grid
	want := [][]Entity{{3, 0, 1, 1, 0, 1, 0, 2}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}
	moves := []Dir{RIGHT}
	w.Step(moves)

	got = w.Grid
	want = [][]Entity{{0, 3, 1, 1, 0, 1, 0, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 0, 3, 1, 1, 1, 0, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 0, 0, 3, 1, 1, 1, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = w.Grid
	want = [][]Entity{{0, 0, 0, 3, 1, 1, 1, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
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

	got := printMap(warehouse.Grid)
	want := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`
	if got != want {
		t.Fatalf("TestMapFull Move \ngot\n%s\nwant\n%s", got, want)
	}

	got2 := computeGPSSum(&warehouse)
	want2 := 10092
	if got2 != want2 {
		t.Fatalf("Move full: got %d want %d", got2, want2)
	}

}

func TestReadMapExtended_mini(t *testing.T) {
	input := `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######`
	lines := strings.Split(input, "\n")

	w := readMapExtended(lines)
	got := printMap(w.Grid)
	want := `##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############`
	if got != want {
		t.Fatalf("readMapExtended \ngot\n%s\nwant\n%s", got, want)
	}
}

func TestReadMapExtended_large(t *testing.T) {
	input := `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########`
	lines := strings.Split(input, "\n")

	w := readMapExtended(lines)
	got := printMap(w.Grid)
	want := `####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##....[]@.....[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################`
	if got != want {
		t.Fatalf("readMapExtended \ngot\n%s\nwant\n%s", got, want)
	}
}

func TestSolvePart1(t *testing.T) {
	got := solvePart1("map.txt", "moves.txt")
	want := 1516281
	if got != want {
		t.Fatalf("SolvePart1 full: got %d want %d", got, want)
	}

}

func TestMoveUpExtended_up(t *testing.T) {
	input := `##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.....@....##
##############`
	lines := strings.Split(input, "\n")
	w := readMap(lines)

	w.Step([]Dir{UP})
	got := printMap(w.Grid)
	want := `##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

	w.Step([]Dir{UP})
	got = printMap(w.Grid)
	want = `##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

	w.Step([]Dir{RIGHT, UP, LEFT})
	got = printMap(w.Grid)
	want = `##############
##......##..##
##...[][]...##
##...[]@....##
##..........##
##..........##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

	w.Step([]Dir{DOWN, LEFT, LEFT, UP, UP, UP})
	got = printMap(w.Grid)
	want = `##############
##...[].##..##
##...[][]...##
##...@......##
##..........##
##..........##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

}

func TestMoveUpExtended_up2(t *testing.T) {
	input := `##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.....@....##
##############`
	lines := strings.Split(input, "\n")
	w := readMap(lines)

	w.Step([]Dir{UP, UP, UP})
	got := printMap(w.Grid)
	want := `##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

	w.Step([]Dir{LEFT, LEFT, UP, UP})
	got = printMap(w.Grid)
	want = `##############
##...[].##..##
##...@.[]...##
##....[]....##
##..........##
##..........##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}
}

func TestMoveUpExtended_down(t *testing.T) {
	input := `##############
##......##..##
##.....@....##
##...[][]...##
##....[]....##
##..........##
##############`
	lines := strings.Split(input, "\n")
	w := readMap(lines)

	w.Step([]Dir{DOWN, DOWN})
	got := printMap(w.Grid)
	want := `##############
##......##..##
##..........##
##...[]@....##
##.....[]...##
##....[]....##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

	w.Step([]Dir{LEFT, LEFT, LEFT, LEFT})
	got = printMap(w.Grid)
	want = `##############
##......##..##
##..........##
##[]@.......##
##.....[]...##
##....[]....##
##############`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}
}

func TestMoveUpExtended_down_special(t *testing.T) {
	input := `##..@.##
##.[].##
##[]..##
##..[]##
##....##
########`
	lines := strings.Split(input, "\n")
	w := readMap(lines)

	w.Step([]Dir{DOWN})
	got := printMap(w.Grid)

	want := `##....##
##..@.##
##.[].##
##[][]##
##....##
########`
	if got != want {
		t.Fatalf("moveUpExtended \ngot\n%s\nwant\n%s", got, want)
	}

}

func TestMove_right_extended(t *testing.T) {
	input1 := `@.[].[].[].#`
	w := readMap([]string{input1})

	got := printMap(w.Grid)
	want := `@.[].[].[].#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves := []Dir{RIGHT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `.@[].[].[].#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `..@[][].[].#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `...@[][][].#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `....@[][][]#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `....@[][][]#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{LEFT, LEFT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `..@..[][][]#`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}
}

func TestMove_left_extended(t *testing.T) {
	input1 := `#.[][]..[].@`
	w := readMap([]string{input1})

	got := printMap(w.Grid)
	want := `#.[][]..[].@`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves := []Dir{LEFT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `#.[][]..[]@.`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{LEFT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `#.[][].[]@..`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{LEFT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `#.[][][]@...`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{LEFT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `#[][][]@....`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{LEFT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `#[][][]@....`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", got, want)
	}

	moves = []Dir{RIGHT, RIGHT}
	w.Step(moves)
	got = printMap(w.Grid)
	want = `#[][][]..@..`
	if got != want {
		t.Errorf("Grid after move right: \ngot  %v\nwant %v", want, got)
	}
}

func TestMoveExtended_full(t *testing.T) {
	rawData, err := os.ReadFile("map_mini.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawData), "\n")
	warehouse := readMapExtended(lines)
	moves, err := readMoves("moves_mini.txt")
	if err != nil {
		panic(err)
	}
	warehouse.Step(moves)

	got := printMap(warehouse.Grid)
	want := `####################
##[].......[].[][]##
##[]...........[].##
##[]........[][][]##
##[]......[]....[]##
##..##......[]....##
##..[]............##
##..@......[].[][]##
##......[][]..[]..##
####################`

	if got != want {
		t.Errorf("Grid after full move extended: \ngot  \n%v\nwant \n%v", got, want)
	}

	got2 := computeGPSSum(&warehouse)
	want2 := 9021
	if got2 != want2 {
		t.Fatalf("Move full: got %d want %d", got2, want2)
	}

}
