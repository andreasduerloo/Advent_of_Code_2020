package day_05

import (
	"os"
	"strings"
	"testing"
)

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
