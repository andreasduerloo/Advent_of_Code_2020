package day_01

import (
	"os"
	"testing"
)

func BenchmarkBinarySearch1(b *testing.B) {
	// Generate the slice of integers
	input, _ := os.ReadFile("../inputs/01.txt")

	expenses := parse(string(input))

	// Reset the timer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, expense := range expenses {
			_ = checkNumber(expense, expenses)
		}
	}
}

func BenchmarkNaive1(b *testing.B) {
	// Generate the slice of integers
	input, _ := os.ReadFile("../inputs/01.txt")

	expenses := parse(string(input))

	// Reset the timer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, expense := range expenses {
			_ = naive1(expense, expenses)
		}
	}
}

func BenchmarkBinarySearch2(b *testing.B) {
	// Generate the slice of integers
	input, _ := os.ReadFile("../inputs/01.txt")

	expenses := parse(string(input))

	// Reset the timer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, expense := range expenses {
			_ = checkThree(expense, expenses)
		}
	}
}

func BenchmarkNaive2(b *testing.B) {
	// Generate the slice of integers
	input, _ := os.ReadFile("../inputs/01.txt")

	expenses := parse(string(input))

	// Reset the timer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, expense := range expenses {
			_ = naive2(expense, expenses)
		}
	}
}
