package day1

import "testing"

func BenchmarkDay1Stream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day1Stream("./input.txt")
	}
}

func BenchmarkDay1Loaded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day1Loaded("./input.txt")
	}
}
