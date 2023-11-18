package day_04

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/04.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	passports := parse(string(input))

	var first, second int
	valid := make([]map[string]string, 0)

	for _, passport := range passports {
		if valid1(passport) {
			first += 1
			valid = append(valid, passport)
		}
	}

	for _, passport := range valid {
		if valid2(passport) {
			second += 1
		}
	}

	return first, second
}
