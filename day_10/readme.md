# Adapter Array

This day's first star was nothing special, but the **second star** is much more interesting. I would argue it is the first problem in this edition where some knowledge of algorithms/dynamic programming is **required** to find a solution that computes in an acceptable amount of time (whereas before it was just a nice extra).

The problem can be expressed pretty naturally through recursion: for any given adapter, the amount of combinations leading up to that adapter is equal to the sum of the amount of combinations leading up to each of the possible adapters immediately before this one.

```go
func combinations(joltages []int) int {
	var localsum int
	var candidates []int

	if len(joltages) >= 4 {                     // Figure out how many adapters are left.
		candidates = joltages[1:4]              // if we don't do this, we will run into
	} else {                                    // out-of-bounds errors once the slices
		candidates = joltages[1:]               // get shorter.
	}

	for i, candidate := range candidates {                  // Only retain those adapters we can connect to this one.
		if joltages[0]-candidate <= 3 {                     // Add the number of combinations for each of the candidates
			localsum += combinations(joltages[i+1:])        // to a local sum.
		}
	}

	return localsum
}
```
Given that we have not defined any base cases, the function above will either run forever or (more likely) run into an error once we start hitting empty slices. There are two base cases to foresee:
- Either we're looking at a slice containing a single adapter: there's one way to reach this, we return 1.
- Either we're looking at two adapters: there's also only one way to reach this, return 1.
From then on we start passing the values back up the recursion tree.

This approach will work in theory, and it will even work in practice for the example inputs, but it will run for a **long** time on the real problem input. Why is that?

As we keep splitting at "junctions" (i.e., more then one adapter could connect to the one we're looking at), we will very quickly start recalculating the same values. In other words: our recursion tree will consist mostly of identical branches or sub-branches. We are doing far too much work.

What we need to do is use **memoization**. This means we keep track of values we have calculated before. If we're about to recalculate one of those values, we look up the result we got the last time instead of doing all the work again. Using this trick, the output is calculated near-instantanously. This combination of recursion and memoization is often called *top-down dynamic programming*.