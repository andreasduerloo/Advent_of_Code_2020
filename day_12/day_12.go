package day_12

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/12.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	instructions := parse(input)

	s := ship{
		direction: 'E',
		x:         0,
		y:         0,
	}

	for _, instruction := range instructions {
		s.apply(instruction)
	}

	first := distance(s)

	return first, 0
}
