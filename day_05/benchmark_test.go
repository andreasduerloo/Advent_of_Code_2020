package day_05

import (
	"os"
	"strings"
	"testing"
)

// Testing the ways of turning the seat codes into integers
func BenchmarkBinarySearch(b *testing.B) { // Three times faster, but doing the entire loop adds so much overhead that it's only 25% faster
	/*
		// Read the input
		input, _ := os.ReadFile("../inputs/05.txt")
		lines := strings.Split(string(input), "\n")

		// Reset the timer
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				if line != "" {
					_ = seatID(line)
				}
			}
		}
	*/
	for i := 0; i < b.N; i++ {
		binary("BFFFFFFRLL")
	}
}

func BenchmarkBinaryConversion(b *testing.B) {
	/*
		// Read the input
		input, _ := os.ReadFile("../inputs/05.txt")
		lines := strings.Split(string(input), "\n")

		// Reset the timer
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				if line != "" {
					_ = binary(line)
				}
			}
		}
	*/
	for i := 0; i < b.N; i++ {
		seatID("BFFFFFFRLL")
	}
}

// Testing the algorithm for finding the missing seat

func BenchmarkGauss(b *testing.B) {
	// Generate the slice of integers
	input, _ := os.ReadFile("../inputs/05.txt")

	lines := strings.Split(string(input), "\n")

	seatIDs := make([]int, 0)

	for _, line := range lines {
		if line != "" {
			id := seatID(line)
			seatIDs = append(seatIDs, id)
		}
	}

	// Reset the timer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		missingSeat(seatIDs)
	}
}

func BenchmarkIter(b *testing.B) {
	// Generate the slice of integers
	input, _ := os.ReadFile("../inputs/05.txt")

	lines := strings.Split(string(input), "\n")

	seatIDs := make([]int, 0)

	for _, line := range lines {
		if line != "" {
			id := seatID(line)
			seatIDs = append(seatIDs, id)
		}
	}

	// Reset the timer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		missingSeatIter(seatIDs)
	}
}
