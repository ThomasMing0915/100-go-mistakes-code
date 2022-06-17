package main

import (
	"fmt"
	"time"
)

func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					fmt.Println("case ch1 close.")
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					fmt.Println("case ch2 close.")
					break
				}
				ch <- v
			}
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

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	close(ch1)
	time.Sleep(time.Second)
	close(ch2)
	time.Sleep(time.Second)

}
