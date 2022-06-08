package main

import (
	"fmt"
	"time"
)

func example1() {
	s := []int{1, 2, 3}
	for _, i := range s {
		go func() {
			fmt.Println(i)
		}()
	}
}

func example2() {
	s := []int{1, 2, 3}
	for _, i := range s {
		val := i
		go func() {
			fmt.Println(val)
		}()
	}
}

func example3() {
	s := []int{1, 2, 3}
	for _, i := range s {
		go func(val int) {
			fmt.Println(val)
		}(i)
	}
}

func main() {
	fmt.Println("example1")
	example1()
	time.Sleep(time.Second)

	fmt.Println("\nexample2")
	example2()
	time.Sleep(time.Second)

	fmt.Println("\nexample3")
	example3()
	time.Sleep(time.Second)
}
