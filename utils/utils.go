package utils

import (
	"fmt"
	"os"
	"strings"
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
