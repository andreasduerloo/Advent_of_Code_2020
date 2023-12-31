package day_11

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/11.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	plan := parse(input)

	for plan.changed {
		plan = plan.apply1()
	}

	first := countOccupied(plan)

	plan = parse(input)

	for plan.changed {
		plan = plan.apply2()
	}

	second := countOccupied(plan)

	return first, second
}
