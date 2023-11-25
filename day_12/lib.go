package day_12

import (
	"advent2020/helpers"
	"strings"
)

type ship struct {
	x         int
	y         int
	direction rune
}

type instruction struct {
	action rune
	number int
}

func (s *ship) apply(instr instruction) {
	switch instr.action {
	case 'N':
		s.y += instr.number
	case 'S':
		s.y -= instr.number
	case 'E':
		s.x += instr.number
	case 'W':
		s.x -= instr.number
	case 'F':
		s.apply(instruction{action: s.direction, number: instr.number})
	default:
		s.direction = turn(s.direction, instr.action, instr.number)
	}
}

func turn(course, lr rune, deg int) rune {
	directions := []rune{'N', 'E', 'S', 'W'}
	var dirIndex int
	turns := deg / 90

	for i, dir := range directions {
		if dir == course {
			dirIndex = i
		}
	}

	switch lr {
	case 'L':
		return directions[(dirIndex+4-turns)%4]
	case 'R':
		return directions[(dirIndex+turns)%4]
	}

	return 'X'
}

func parse(input []byte) []instruction {
	lines := strings.Split(string(input), "\n")
	outSlice := make([]instruction, 0)

	for _, line := range lines {
		if line != "" {
			direction := line[0]
			num := helpers.ReGetInts(line)[0]

			outSlice = append(outSlice, instruction{action: rune(direction), number: num})
		}
	}
	return outSlice
}

func distance(s ship) int {
	var x, y int

	if s.x < 0 {
		x = s.x * -1
	} else {
		x = s.x
	}

	if s.y < 0 {
		y = s.y * -1
	} else {
		y = s.y
	}

	return x + y
}
