package main

import "testing"

func BenchmarkParallelMergesortV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parallelMergesortV2(sV2)
	}
}
