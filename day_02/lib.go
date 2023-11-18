package day_02

import (
	"advent2020/helpers"
	"strings"
)

type prompt struct {
	low      int
	high     int
	letter   rune
	password string
}

func parse(input string) []*prompt {
	lines := strings.Split(input, "\n")
	out := make([]*prompt, 0)

	for _, line := range lines {
		if line != "" {
			ints := helpers.ReGetInts(line)
			fields := strings.Fields(line)
			out = append(out, &prompt{
				low:      ints[0],
				high:     ints[1],
				letter:   rune(fields[1][0]),
				password: fields[2],
			})
		}
	}
	return out
}

func (p *prompt) validate1() bool {
	tocount := p.letter
	var count int

	for _, r := range p.password {
		if r == tocount {
			count += 1
		}
	}

	if p.low <= count && count <= p.high {
		return true
	}
	return false
}

func (p *prompt) validate2() bool {
	var first, second bool

	first = p.letter == rune(p.password[p.low-1])
	second = p.letter == rune(p.password[p.high-1])

	if first != second { // Nifty XOR, alternative to first || second && !(first && second)
		return true
	}
	return false
}
