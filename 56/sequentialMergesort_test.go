package main

import "testing"

func BenchmarkSequentialMergesort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sequentialMergesort(sSeq)
	}
}
