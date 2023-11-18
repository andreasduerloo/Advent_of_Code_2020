package day_06

import (
	"strings"
)

func parse(input string) [][]string {
	out := make([][]string, 0)
	groups := strings.Split(input, "\n\n")

	for _, group := range groups {
		out = append(out, strings.Split(group, "\n"))
	}

	return out
}

func groupScore(group []string) (int, int) {
	set := make(map[rune]int)
	var out1, out2 int

	for _, form := range group {
		for _, question := range form { // First person to answer yes?
			if _, present := set[question]; !present {
				out1 += 1
				set[question] = 1
			} else {
				set[question] = set[question] + 1
			}
		}
	}

	for _, v := range set {
		if v == len(group) {
			out2 += 1
		}
	}

	return out1, out2
}
