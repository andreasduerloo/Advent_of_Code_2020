# Day 05 - Binary Boarding

For this puzzle I found a solution for the second star that *feels* smart, but might not be. Benchmarking that solution against a more naive one could clear it up.

Here's the problem: you receive a list of contiguous integers (not necessarily starting at zero or one), with **one** integer missing. How do you find that integer?

Here's my naive solution:
```go
func missingSeatIter(ids []int) int {
	slices.Sort(ids)

	for i, id := range ids {
		if ids[i+1]-id > 1 {
			return id + 1
		}
	}
	return 0
}
```
First I sort all the integers in increasing order, and then I iterate through them one-by-one, always comparing an integer with the next one. If ever the difference between those two integers is greater than 1, I return the integer in between. This gives the correct answer. This approach has two downsides:
1. I have to sort the entire array - Go uses quicksort ([source](https://cs.opensource.google/go/go/+/master:src/slices/zsortordered.go;l=63?q=pdqsortOrdered&sq=&ss=go%2Fgo)), which has a time complexity of O(n log n).
2. I have to iterate through (potentially) the entire array

On the other hand, here is my "smart" solution:
```go
func missingSeat(ids []int) int {
	lowest := slices.Min(ids)
	highest := slices.Max(ids)

	expectedSum := (lowest + highest) * ((highest - lowest) / 2)

	if (len(ids)+1)%2 == 1 { // Add one because we know the list is one short
		expectedSum += (highest + lowest) / 2
	}

	var actualSum int
	for _, id := range ids {
		actualSum += id
	}

	return expectedSum - actualSum
}
```
What am I doing here? I note the lowest and highest integers. Then, using Gauss' trick for summing a series of integers, I calculate what we expect the sum of all integers to be if we didn't have one missing. Then we calculate the actual sum and subtract the latter from the former. This also gives the correct answer. This approach also has a few drawbacks:
1. The Min and Max functions in the slice package iterate through all integers (not that there's another way), [see source.](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/slices/sort.go;l=61) That means I'm scanning through the entire thing twice already. This could be solved by writing my own function that collects both the lowest and highest int in **one** scan.
2. Then I'm summing all of the integers, so I **know** I'm iterating through all of the integers another time.
3. I can't exit out of the summing early, as opposed to the first approach where the missing integer could potentially be found very early on.

I benchmarked these two functions, and there's the output:
```
BenchmarkGauss-4          498862              2033 ns/op
BenchmarkIter-4           902784              1407 ns/op
```
That would mean the Gauss-version is almost 50% slower than the (naive) iterative one, at least for this input size. Let's see what happens if I replace the calls to `slices.Min()` and `slices.Max()` with a single function:
```go
func minMax(ids []int) (int, int) {
	min, max := ids[0], ids[0]
	for _, id := range ids {
		if id < min {
			min = id
		}
		if id > max {
			max = id
		}
	}

	return min, max
}

func missingSeat(ids []int) int {
	lowest, highest := minMax(ids)

	expectedSum := (lowest + highest) * ((highest - lowest) / 2)

	if (len(ids)+1)%2 == 1 { // Add one because we know the list is one short
		expectedSum += (highest + lowest) / 2
	}

	var actualSum int
	for _, id := range ids {
		actualSum += id
	}

	return expectedSum - actualSum
}
```
I run the benchmark again:
```
BenchmarkGauss-4          970496              1264 ns/op
BenchmarkIter-4           764760              1565 ns/op
```
Now the Gauss implementation is faster than the iterative one! That would mean that for this input size, getting rid of one linear scan makes all the difference.

However, we still have another linear scan going on when we are calculating the actual sum. Let's make that part of the same loop as well:
```go
func minMaxSum(ids []int) (int, int, int) {
	min, max, sum := ids[0], ids[0], 0
	for _, id := range ids {
		sum += id
		if id < min {
			min = id
		}
		if id > max {
			max = id
		}
	}

	return min, max, sum
}

func missingSeat(ids []int) int {
	lowest, highest, actualSum := minMaxSum(ids)

	expectedSum := (lowest + highest) * ((highest - lowest) / 2)

	if (len(ids)+1)%2 == 1 { // Add one because we know the list is one short
		expectedSum += (highest + lowest) / 2
	}

	return expectedSum - actualSum
}
```
New benchmark:
```
BenchmarkGauss-4         1249065               971.7 ns/op
BenchmarkIter-4           913374              1315 ns/op
```
Faster again! And, more importantly: twice as fast as the initial implementation!