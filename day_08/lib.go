package day_08

import (
	"strconv"
	"strings"
)

func runUntilLoop(input []string) (int, bool) {
	var pos, acc int
	history := make(map[int]struct{})

	_, present := history[pos]
	for !present { // Check for presence before executing
		history[pos] = struct{}{}
		fields := strings.Fields(input[pos])
		switch fields[0] {
		case "acc":
			num, _ := strconv.Atoi(fields[1])
			acc += num
			pos++
		case "jmp":
			num, _ := strconv.Atoi(fields[1])
			pos += num
		case "nop":
			pos++
		}
		_, present = history[pos]
		if pos >= len(input)-1 {
			return acc, false
		}
	}

	return acc, true // We're in a loop
}

func runUntilTerminates(instructions []string) int {

	for i := 0; i < len(instructions); i++ {
		fields := strings.Fields(instructions[i])
		var acc int
		loop := false

		switch fields[0] {
		case "jmp":
			instructions[i] = "nop " + fields[1]
			acc, loop = runUntilLoop(instructions)
			// Put it back the way it was
			instructions[i] = "jmp " + fields[1]
		case "nop":
			instructions[i] = "jmp " + fields[1]
			acc, loop = runUntilLoop(instructions)
			// Put it back the way it was
			instructions[i] = "nop " + fields[1]
		default:
			continue
		}
		if !loop {
			return acc
		}
		acc = 0
	}
	return 0
}
