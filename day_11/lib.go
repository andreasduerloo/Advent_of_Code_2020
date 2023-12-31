package day_11

import (
	"strings"
)

type grid struct {
	content []rune
	width   int
	height  int
	changed bool
}

func parse(input []byte) grid {
	var width, height int
	outSlice := make([]rune, 0)

	lines := strings.Split(string(input), "\n")
	width = len(lines[0])
	height = len(lines) - 1

	for _, line := range lines {
		for _, char := range line {
			outSlice = append(outSlice, rune(char))
		}
	}

	return grid{
		content: outSlice,
		width:   width,
		height:  height,
		changed: true,
	}
}

func (old grid) apply1() grid {
	outArr := make([]rune, 0)
	var changed bool

	for i, char := range old.content {
		if char == '.' {
			outArr = append(outArr, '.')
		} else {
			var occupied int
			if i >= old.width && i%old.width != 0 { // Up and left
				if old.content[i-old.width-1] == '#' {
					occupied++
				}
			}
			if i >= old.width && i%old.width != old.width-1 { // Up and right
				if old.content[i-old.width+1] == '#' {
					occupied++
				}
			}
			if i < (old.width*(old.height-1)) && i%old.width != 0 { // Down and left
				if old.content[i+old.width-1] == '#' {
					occupied++
				}
			}
			if i < (old.width*(old.height-1)) && i%old.width != old.width-1 { // Down and right
				if old.content[i+old.width+1] == '#' {
					occupied++
				}
			}
			if i%old.width != 0 { // Left edge?
				if old.content[i-1] == '#' {
					occupied++
				}
			}
			if i%old.width != old.width-1 { // Right edge?
				if old.content[i+1] == '#' {
					occupied++
				}
			}
			if i >= old.width { // Top edge?
				if old.content[i-old.width] == '#' {
					occupied++
				}
			}
			if i < (old.width*old.height)-old.width { // Bottom edge?
				if old.content[i+old.width] == '#' {
					occupied++
				}
			}

			if char == '#' && occupied >= 4 {
				outArr = append(outArr, 'L')
				changed = true
			} else if char == '#' && occupied < 4 {
				outArr = append(outArr, '#')
			} else if char == 'L' && occupied == 0 {
				outArr = append(outArr, '#')
				changed = true
			} else if char == 'L' && occupied > 0 {
				outArr = append(outArr, 'L')
			}
		}

	}

	return grid{
		content: outArr,
		width:   old.width,
		height:  old.height,
		changed: changed,
	}
}

func (old grid) apply2() grid {
	outArr := make([]rune, 0)
	var changed bool

	for j, char := range old.content {
		if char == '.' {
			outArr = append(outArr, '.')
		} else {
			var occupied int
			i := j
			for i >= old.width && i%old.width != 0 { // Look up and left
				if old.content[i-old.width-1] == '#' {
					occupied++
					break
				} else if old.content[i-old.width-1] == 'L' {
					break
				} else {
					i = i - old.width - 1
				}
			}
			i = j
			for i >= old.width && i%old.width != old.width-1 { // Look up and right
				if old.content[i-old.width+1] == '#' {
					occupied++
					break
				} else if old.content[i-old.width+1] == 'L' {
					break
				} else {
					i = i - old.width + 1
				}
			}
			i = j
			for i < (old.width*(old.height-1)) && i%old.width != 0 { // Look down and left
				if old.content[i+old.width-1] == '#' {
					occupied++
					break
				} else if old.content[i+old.width-1] == 'L' {
					break
				} else {
					i = i + old.width - 1
				}
			}
			i = j
			for i < (old.width*(old.height-1)) && i%old.width != old.width-1 { // Look down and right
				if old.content[i+old.width+1] == '#' {
					occupied++
					break
				} else if old.content[i+old.width+1] == 'L' {
					break
				} else {
					i = i + old.width + 1
				}
			}
			i = j
			for i%old.width != 0 { // Look left
				if old.content[i-1] == '#' {
					occupied++
					break
				} else if old.content[i-1] == 'L' {
					break
				} else {
					i--
				}
			}
			i = j
			for i%old.width != old.width-1 { // Look right
				if old.content[i+1] == '#' {
					occupied++
					break
				} else if old.content[i+1] == 'L' {
					break
				} else {
					i++
				}
			}
			i = j
			for i >= old.width { // Look up
				if old.content[i-old.width] == '#' {
					occupied++
					break
				} else if old.content[i-old.width] == 'L' {
					break
				} else {
					i = i - old.width
				}
			}
			i = j
			for i < (old.width*old.height)-old.width { // Look down
				if old.content[i+old.width] == '#' {
					occupied++
					break
				} else if old.content[i+old.width] == 'L' {
					break
				} else {
					i = i + old.width
				}
			}

			if char == '#' && occupied >= 5 {
				outArr = append(outArr, 'L')
				changed = true
			} else if char == '#' && occupied < 5 {
				outArr = append(outArr, '#')
			} else if char == 'L' && occupied == 0 {
				outArr = append(outArr, '#')
				changed = true
			} else if char == 'L' && occupied > 0 {
				outArr = append(outArr, 'L')
			}
		}
	}

	return grid{
		content: outArr,
		width:   old.width,
		height:  old.height,
		changed: changed,
	}
}

func countOccupied(seats grid) int {
	var out int

	for _, seat := range seats.content {
		if seat == '#' {
			out++
		}
	}

	return out
}
