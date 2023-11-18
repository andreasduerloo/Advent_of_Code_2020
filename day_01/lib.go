package day_01

import (
	"slices"
	"strconv"
	"strings"
)

func parse(s string) []int {
	out := make([]int, 0)
	vals := strings.Split(s, "\n")

	for _, v := range vals {
		intval, _ := strconv.Atoi(v)
		out = append(out, intval)
	}

	slices.Sort(out)

	return out
}

func binarySearch(num int, all []int) bool {
	index := len(all) / 2

	for len(all) > 1 {
		if num < all[index] {
			all = all[:index]
			index = len(all) / 2
		} else if num > all[index] {
			all = all[index:]
			index = len(all) / 2
		} else {
			return true
		}
	}

	if all[0] == num {
		return true
	}

	return false
}

func checkNumber(num int, all []int) int {
	needed := 2020 - num

	if binarySearch(needed, all) {
		return num * needed
	}

	return 0
}

func checkThree(num int, all []int) int {
	for _, other := range all {
		if (num + other) < 2020 {
			needed := 2020 - num - other

			if binarySearch(needed, all) {
				return needed * num * other
			}
		}
	}
	return 0
}

// For benchmarking

func naive1(num int, all []int) bool {
	for _, expense := range all {
		for _, other := range all {
			if 2020 == expense+other {
				return true
			}
		}
	}
	return false
}

func naive2(num int, all []int) bool {
	for _, expense := range all {
		for _, second := range all {
			for _, third := range all {
				if 2020 == expense+second+third {
					return true
				}
			}
		}
	}
	return false
}
