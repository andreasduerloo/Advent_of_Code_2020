package day_09

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/09.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	numbers := parse(input)

	first := firstIncorrect(numbers)

	return first, 0
}
