package day_10

import (
	"slices"
	"strconv"
	"strings"
)

func parse(input []byte) []int { // Second time input works like this -> add to helpers?
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

func jumps(joltages []int) int {
	var one, three int

	slices.Sort(joltages)

	if joltages[0] == 1 { // Initial jump from 0
		one++
	}

	if joltages[0] == 3 {
		three++
	}

	for i := 0; i < len(joltages)-1; i++ {
		switch joltages[i+1] - joltages[i] {
		case 1:
			one++
		case 3:
			three++
		default:
			continue
		}
	}

	return one * (three + 1)
}

func combinations(joltages []int, setup bool, cache map[int]int) int {
	if setup {
		joltages = append(joltages, joltages[len(joltages)-1]+3)
		slices.Reverse(joltages)
		joltages = append(joltages, 0)
	}

	if len(joltages) == 1 {
		return 1
	}

	if len(joltages) == 2 {
		if joltages[0]-joltages[1] <= 3 {
			return 1
		} else {
			return 0
		}
	}

	var localsum int
	var candidates []int
	if len(joltages) >= 4 {
		candidates = joltages[1:4]
	} else {
		candidates = joltages[1:]
	}

	for i, candidate := range candidates {
		if joltages[0]-candidate <= 3 {
			if val, present := cache[candidate]; present {
				localsum += val
			} else {
				localsum += combinations(joltages[i+1:], false, cache)
			}
		}
	}

	if localsum != 0 {
		cache[joltages[0]] = localsum
	}
	return localsum
}

// Slice prepend: https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice
// joltages = append([]int{0}, joltages...)
