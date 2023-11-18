package day_05

import (
	"slices"
)

func bsp(s string, low, high int) int { // Recursive
	if len(s) == 1 {
		switch s[0] {
		case 'F', 'L':
			return low
		case 'B', 'R':
			return high
		default:
			return 0
		}
	}

	switch s[0] {
	case 'F', 'L':
		return bsp(s[1:], low, (low+high)/2)
	case 'B', 'R':
		return bsp(s[1:], (low+high+1)/2, high)
	default:
		return 0
	}
}

func binary(s string) int { // Turn the string directly into the binary number it will turn out to be, anyway
	var out int

	for i, r := range s {
		if r == 'B' || r == 'R' {
			out += (1 << (9 - i))
		}
	}

	return out
}

func seatID(seat string) int {
	return (8 * bsp(seat[0:7], 0, 127)) + bsp(seat[7:], 0, 7)
}

func missingSeat(ids []int) int {
	lowest, highest, actualSum := minMaxSum(ids)

	expectedSum := (lowest + highest) * ((highest - lowest) / 2)

	if (len(ids)+1)%2 == 1 { // Add one because we know the list is one short
		expectedSum += (highest + lowest) / 2
	}

	return expectedSum - actualSum
}

func minMaxSum(ids []int) (int, int, int) {
	min, max, sum := ids[0], ids[0], 0
	for _, id := range ids {
		sum += id
		if id < min {
			min = id
		}
		if id > max {
			max = id
		}
	}

	return min, max, sum
}

func missingSeatIter(ids []int) int {
	slices.Sort(ids)

	for i, id := range ids {
		if ids[i+1]-id > 1 {
			return id + 1
		}
	}
	return 0
}
