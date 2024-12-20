package main

import (
	"testing"
)

// FunctionA is a hypothetical function that we want to benchmark
func FunctionA() {
	// ... some code here ...
}

// BenchA is the benchmark function for FunctionA
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FunctionA()
	}
}