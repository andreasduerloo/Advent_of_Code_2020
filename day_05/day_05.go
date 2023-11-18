package day_05

import (
	"fmt"
	"os"
	"strings"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/05.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	lines := strings.Split(string(input), "\n")

	var first, second int
	seatIDs := make([]int, 0)

	for _, line := range lines {
		if line != "" {
			id := seatID(line)
			seatIDs = append(seatIDs, id)
			if id > first {
				first = id
			}
		}
	}

	second = missingSeat(seatIDs)

	return first, second
}
