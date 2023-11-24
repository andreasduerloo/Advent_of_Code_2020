package day_10

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/10.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	joltages := parse(input)

	first := jumps(joltages)

	cache := make(map[int]int)
	second := combinations(joltages, true, cache)

	return first, second
}
