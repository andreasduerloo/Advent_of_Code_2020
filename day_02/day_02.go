package day_02

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/02.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	prompts := parse(string(input))
	var first, second int

	for _, p := range prompts {
		if p.validate1() {
			first += 1
		}
		if p.validate2() {
			second += 1
		}
	}

	return first, second
}
