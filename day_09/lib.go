package day_09

import (
	"advent2020/helpers"
	"strconv"
	"strings"
)

func parse(input []byte) []int {
	out := make([]int, 0)
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line != "" {
			val, err := strconv.Atoi(line)
			if err != nil {
				continue
			}

			out = append(out, val)
		}
	}
	return out
}

func firstIncorrect(nums []int) int {
	queue := make([]int, 0)

	for _, num := range nums {
		if len(queue) < 25 {
			queue = append(queue, num)
		} else {
			if !sumPresent(num, queue) {
				return num
			}

			_, queue = helpers.Dequeue(queue)
			queue = append(queue, num)
		}
	}

	return 0
}

func sumPresent(num int, preamble []int) bool {
	for _, first := range preamble {
		for _, second := range preamble {
			if first != second && first+second == num {
				return true
			}
		}
	}

	return false
}
