package day_06

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/06.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	groups := parse(string(input))

	var first, second int

	for _, group := range groups {
		f, s := groupScore(group)
		first += f
		second += s
	}

	return first, second
}
