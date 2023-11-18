package day_01

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/01.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	expenses := parse(string(input))
	var first, second int

	for _, expense := range expenses {
		first = checkNumber(expense, expenses)
		if first != 0 {
			break
		}
	}

	for _, expense := range expenses {
		second = checkThree(expense, expenses)
		if second != 0 {
			break
		}
	}

	return first, second
}
