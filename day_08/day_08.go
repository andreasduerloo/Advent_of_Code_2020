package day_08

import (
	"fmt"
	"os"
	"strings"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/08.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	instructions := strings.Split(string(input), "\n")
	first, _ := runUntilLoop(instructions)
	second := runUntilTerminates(instructions)

	return first, second
}
