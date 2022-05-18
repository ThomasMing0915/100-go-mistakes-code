package main

func sequentialMergesort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergesort(s[:middle])
	sequentialMergesort(s[middle:])
	merge(s, middle)
}
