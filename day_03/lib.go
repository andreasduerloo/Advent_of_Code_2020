package day_03

import "strings"

func parse(s string) []string {
	out := strings.Split(s, "\n")

	if out[len(out)-1] == "" {
		return out[:len(out)-1]
	}

	return out
}

func slope(grid []string, right, down int) int {
	var out, col int
	for i := 0; i < len(grid); i += down {
		if grid[i][col] == '#' {
			out += 1
		}
		col = (col + right) % len(grid[i])
	}

	return out
}
