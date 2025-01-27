package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func PrintSolution[E any](day, part int, res E) {
	fmt.Printf("Day %d part %d solution: %v\n", day, part, res)
}

func ReadLines(path string) ([]string, error) {
	rawStr, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(rawStr), "\n"), nil
}

func Inside(pos [2]int, rows, cols int) bool {
	if pos[1] < 0 || pos[1] >= cols {
		return false
	}
	if pos[0] < 0 || pos[0] >= rows {
		return false
	}
	return true
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
