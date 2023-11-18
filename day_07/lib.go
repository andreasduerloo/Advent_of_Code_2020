package day_07

import (
	"regexp"
	"strconv"
	"strings"
)

type color struct {
	name       string
	canContain map[string]int
}

func parse(input string) map[string]*color {
	out := make(map[string]*color)

	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			newCol := newColor(line)
			out[newCol.name] = newCol
		}
	}

	return out
}

func newColor(s string) *color {
	fields := strings.Fields(s)
	out := color{
		name:       fields[0] + " " + fields[1], // Could also work with a regex with an anchor: `^[a-z]+ [a-z]+`, that would grab the first two words.
		canContain: children(s),
	}

	return &out
}

func children(s string) map[string]int {
	out := make(map[string]int, 0)

	re := regexp.MustCompile(`\d[a-z ]+[,.]`)
	numberRe := regexp.MustCompile(`\d+`)
	colorRe := regexp.MustCompile(`[a-z]+ [a-z]+`) // We cut off ' bags[,.]'

	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		color := colorRe.FindString(match)
		number, _ := strconv.Atoi(numberRe.FindString(match))
		out[color] = number
	}

	return out
}

// Start with a map containing just "shiny gold"
func checkColor(c *color, set map[string]struct{}) {
	if _, present := set[c.name]; present { // This color is already in the set, we know it can contain "shiny gold" (in)directly.
		return
	}

	for color, _ := range c.canContain {
		if _, present := set[color]; present { // This color can contain "shiny gold" (in)directly
			set[c.name] = struct{}{} // Add the parent item, because it can contain a useful color
		}
	}
}

func countBags(c *color, bags map[string]*color) int {
	var count int
	if len(c.canContain) == 0 {
		return count // Will be zero
	}

	for col, num := range c.canContain {
		count += num*(countBags(bags[col], bags)) + num // The bags themselves, plus what they contain
	}

	return count
}

/*
type unionFind struct {
	name   string
	parent *unionFind
}

func (u *unionFind) find() *unionFind {
	current := u
	for current.parent != nil {
		current = current.parent
	}

	return current
}

func union(u1, u2 *unionFind) {
	root1, root2 := u1.find(), u2.find()

	if root1 == root2 { // Same set
		return
	}

	// If any of the parents have 'shiny gold' as their root, attach to that tree and adapt the size

	// If not, attach the shorter tree to the longer one and adapt the size
}
*/
