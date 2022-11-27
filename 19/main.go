package main

import "fmt"

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}

func main() {
	fmt.Println(f1(100))
	fmt.Println(f2(100))

	a := 100000.001
	b := 1.0001
	c := 1.0002

	fmt.Println(a * (b + c))
	fmt.Println(a*b + a*c)
}
