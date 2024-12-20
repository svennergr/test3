package main

import (
	"testing"
)

// FunctionA is a hypothetical function that we want to benchmark jsut to be contributor
func FunctionA() {
	test
}

// BenchA is the benchmark function for FunctionA
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FunctionA()
	}
}
