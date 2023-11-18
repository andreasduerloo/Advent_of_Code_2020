package day_07

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/07.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	colors := parse(string(input))

	var noChange bool
	set := map[string]struct{}{"shiny gold": struct{}{}}

	for !noChange {
		setLen := len(set)

		for _, color := range colors {
			checkColor(color, set)
		}

		if len(set) == setLen {
			noChange = true
		}
	}

	first := len(set) - 1 // Subtract shiny gold

	second := countBags(colors["shiny gold"], colors)

	return first, second
}
