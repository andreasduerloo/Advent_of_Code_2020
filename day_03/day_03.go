package day_03

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/03.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	grid := parse(string(input))

	first := slope(grid, 3, 1)
	second := (first * slope(grid, 1, 1) * slope(grid, 5, 1) * slope(grid, 7, 1) * slope(grid, 1, 2))

	return first, second
}
