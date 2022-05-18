package main

import (
	"math/rand"
	"testing"
)

var (
	sSeq []int
	sV1  []int
	sV2  []int
)

const (
	sSize = 10000
)

func init() {
	sSeq = make([]int, sSize)
	for i := 0; i < len(sSeq); i++ {
		sSeq[i] = rand.Intn(100000)
	}

	sV1 = make([]int, sSize)
	sV2 = make([]int, sSize)

	copy(sV1, sSeq)
	copy(sV2, sSeq)
}

func BenchmarkParallelMergesortV1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parallelMergesortV1(sV1)
	}
}
