package main

import "testing"

func Test_sequentialMergesort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sequentialMergesort(tt.args.s)
		})
	}
}
