# Day 07 - Handy Haversacks

Reading the first problem and the input, I can think of two data structure that seem to fit the bill:

- Union-find (because of the first question)
- A directed, weighted graph (because the input mentions how many bags another bag can contain, I suspect the second question will need me to find a minimum or maximum number of bags).

I Initially started implementing a union-find data structure, but the fact that a given color can be contained by multiple other colors (i.e., colors don't have just one parent) makes this approach appear more complicated. Additionally, I have to make sure that 'shiny gold' is the root of one of the sets and make sure I don't accidentally merge it into an other set. I'll come back to this approach later.

First I'll try something more accessible:
1. Make a function that takes a color and adds it to a set if it can contain a color already in the set. Thus the set represents all colors that can contain, directly or indirectly, shiny gold. Shiny gold itself is also a member of that set.
2. Initialize a set containing only 'shiny gold' and scan the input.
3. Keep scanning and adding items to the set (unless the item is already in the set), until the set no longer gets bigger.
4. Read the length of the set.

That worked fine for the first star. I then had to make some changes for the second star:
- Work with maps rather than slices.
- Store how many bags of a given color can be contained, on top of the colors themselves.

With those changes the answer can be found through a recursive function:
```go
func countBags(c *color, bags map[string]*color) int {
	var count int
	if len(c.canContain) == 0 { // This bag doesn't contain any others
		return count // Will be zero
	}

	for col, num := range c.canContain {
		count += num*(countBags(bags[col], bags)) + num // The bags themselves, plus what they contain
	}

	return count
}
```
Going back to my initial two ideas, it seems my current approach is far simpler than setting up a Union-Find data structure, while essentially doing the same thing. Additionally, a weighted, directional graph would have been overkill for the second star. Sometimes less really is more. Nevertheless, it might be a good exercise that I can come back to.