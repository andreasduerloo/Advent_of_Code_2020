package main

import (
	"advent2020/day_01"
	"advent2020/day_02"
	"advent2020/day_03"
	"advent2020/day_04"
	"advent2020/day_05"
	"advent2020/day_06"
	"advent2020/day_07"
	"advent2020/day_08"
	"advent2020/day_09"
	"advent2020/day_10"
	"advent2020/day_11"
	"advent2020/day_12"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument was passed - exiting.")
		return
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The argument is not an integer - exiting.")
	}

	fmt.Println("Solutions for day", day)

	switch day {
	case 1:
		first, second := day_01.Solve()
		fmt.Println(first, second)
	case 2:
		first, second := day_02.Solve()
		fmt.Println(first, second)
	case 3:
		first, second := day_03.Solve()
		fmt.Println(first, second)
	case 4:
		first, second := day_04.Solve()
		fmt.Println(first, second)
	case 5:
		first, second := day_05.Solve()
		fmt.Println(first, second)
	case 6:
		first, second := day_06.Solve()
		fmt.Println(first, second)
	case 7:
		first, second := day_07.Solve()
		fmt.Println(first, second)
	case 8:
		first, second := day_08.Solve()
		fmt.Println(first, second)
	case 9:
		first, second := day_09.Solve()
		fmt.Println(first, second)
	case 10:
		first, second := day_10.Solve()
		fmt.Println(first, second)
	case 11:
		first, second := day_11.Solve()
		fmt.Println(first, second)
	case 12:
		first, second := day_12.Solve()
		fmt.Println(first, second)
	default:
		fmt.Println("That's either not a valid day, or it has not been solved (yet!)")
	}
}
