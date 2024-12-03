package utils

import "fmt"

func PrintSolution[E any](day, part int, res E) {
	fmt.Printf("Day %d part %d solution: %v\n", day, part, res)
}
