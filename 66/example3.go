package main

import "fmt"

func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for v := range ch1 {
			ch <- v
		}
		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func main() {
	ch1, ch2 := make(chan int, 2), make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
	ch2 <- 4

	ch := merge(ch1, ch2)

	for v := range ch {
		fmt.Println(v)
	}
}
