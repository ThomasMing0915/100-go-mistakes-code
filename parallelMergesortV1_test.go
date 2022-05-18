package main

import "testing"

func Test_parallelMergesortV1(t *testing.T) {
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
			parallelMergesortV1(tt.args.s)
		})
	}
}
