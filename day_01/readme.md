# Day 01 - Report Repair

The first time I solved this puzzle (a few years ago, but at most three), I used a naive solution with a lot of nested loops. For the first star I would loop through all numbers and compare each of them to all other numbers (len(input)² operations at most). For the second star I added another level of inefficiency: add each number to each other number, and add every other number to that sum, hoping to find 2020 (so in the worst case, almost the **cube** of len(input) operations).

This time around I have found a more efficient solution. The idea and the outer loop have remained pretty much unchanged: I'm still trying every number. Only this time I have ordered the inputs and added a `binarySearch()` function that can quickly (log(len(input))) check whether the value we need is present in the input. See: [Binary Search Algorithm (Wikipedia)](https://en.wikipedia.org/wiki/Binary_search_algorithm) For the second star I still have a nested loop, but overall the solution is far more streamlined: not only do I re-use the binary search function in the inner loop, but I only check for a third number if the sum of the first two by themselves isn't already 2020 or more (which is the case for **a lot** of the input).

I wrote up a naive solution for both stars as well (consisting of two and three nested loops, respectively), and benchmarked them against the binary search solution. Not breaking when we find the solution (to simulate the worst possible case), we get:
```
BenchmarkBinarySearch1-4          571462              2140 ns/op
BenchmarkNaive1-4                  12013            100041 ns/op
BenchmarkBinarySearch2-4           38438             31548 ns/op
BenchmarkNaive2-4                   8784            141187 ns/op
```
What does this mean?
- For the first star, the binary search solution is almost 50 times more efficient than two nested loops.
- For the second star, the Binary search + nested loop solution is almost five times quicker, but it is fifteen times slower than the first binary search solution, while the naive solution is only 40% slower than the first naive solution.