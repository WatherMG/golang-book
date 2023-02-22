package main

import "testing"

var (
	testData = []string{"a", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c"}
)

func BenchmarkUniq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unique(testData)
	}
}
