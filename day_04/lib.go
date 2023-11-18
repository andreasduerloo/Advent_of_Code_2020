package day_04

import (
	"advent2020/helpers"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func parse(input string) []map[string]string {
	out := make([]map[string]string, 0)
	passports := strings.Split(input, "\n\n")

	for _, pp := range passports {
		if pp != "" {
			fields := strings.Fields(pp)
			ppmap := make(map[string]string)
			for _, field := range fields {
				data := strings.Split(field, ":")
				ppmap[data[0]] = data[1]
			}
			out = append(out, ppmap)
		}
	}

	return out
}

func valid1(passport map[string]string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	out := true

	for _, field := range fields {
		_, present := passport[field]
		out = out && present // Stays true until one field is not present
	}

	return out
}

func valid2(passport map[string]string) bool {
	out := true

	// Birth year
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return false
	}
	out = out && (byr >= 1920 && byr <= 2002)

	if !out {
		return false
	}

	// Issue year
	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		return false
	}
	out = out && (iyr >= 2010 && iyr <= 2020)

	if !out {
		return false
	}

	// Expiration year
	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		return false
	}
	out = out && (eyr >= 2020 && eyr <= 2030)

	if !out {
		return false
	}

	// Height
	height := helpers.ReGetInts(passport["hgt"])[0]
	switch passport["hgt"][len(passport["hgt"])-2:] {
	case "cm":
		out = out && (height >= 150 && height <= 193)
	case "in":
		out = out && (height >= 59 && height <= 76)
	default:
		return false
	}

	if !out {
		return false
	}

	// Hair color
	regex := regexp.MustCompile(`#[[:xdigit:]]{6}`)
	present := regex.Match([]byte(passport["hcl"]))
	out = out && present

	if !out {
		return false
	}

	// Eye color
	valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	out = out && slices.Contains(valid, passport["ecl"])

	if !out {
		return false
	}

	// Passport id
	regex = regexp.MustCompile(`^\d{9}$`)
	out = out && regex.Match([]byte(passport["pid"]))

	return out
}
